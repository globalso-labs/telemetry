/*
 * telemetry
 * provider.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 9 Jul 2024 01:53:20 -0500 by nick.
 *
 * DISCLAIMER: This software is provided "as is" without warranty of any kind, either expressed or implied. The entire
 * risk as to the quality and performance of the software is with you. In no event will the author be liable for any
 * damages, including any general, special, incidental, or consequential damages arising out of the use or inability
 * to use the software (that includes, but not limited to, loss of data, data being rendered inaccurate, or losses
 * sustained by you or third parties, or a failure of the software to operate with any other programs), even if the
 * author has been advised of the possibility of such damages.
 * If a license file is provided with this software, all use of this software is governed by the terms and conditions
 * set forth in that license file. If no license file is provided, no rights are granted to use, modify, distribute,
 * or otherwise exploit this software.
 */

package logger

import (
	"context"
	"fmt"

	"go.globalso.dev/x/telemetry/internal"
	"go.globalso.dev/x/telemetry/internal/constants"
	"go.globalso.dev/x/telemetry/pkg/errors"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/sdk/log"
)

type Holder struct {
	provider  *log.LoggerProvider
	processor log.Processor
	exporter  log.Exporter
}

func NewLogger(ctx context.Context, config *Options) (*Holder, error) {
	if !config.IsEnabled() {
		return nil, errors.ErrTelemetryLogsNotEnabled
	}

	holder := new(Holder)

	// Create the exporter.
	exporter, err := newExporter(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create log exporter: %w", err)
	}
	holder.exporter = exporter

	// Create the processor.
	holder.processor = newProcessor(exporter)

	// Create the holder provider.
	holder.provider = newLoggerProvider(holder.processor)

	return holder, nil
}

func (l *Holder) Provider() *log.LoggerProvider {
	return l.provider
}

func (l *Holder) Shutdown(ctx context.Context) error {
	if err := l.provider.Shutdown(ctx); err != nil {
		return err
	}

	if err := l.processor.Shutdown(ctx); err != nil {
		return err
	}

	if err := l.exporter.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (l *Holder) Close() error {
	ctx := context.Background()
	return _handler.Shutdown(ctx)
}

func newExporter(ctx context.Context, _ *Options) (*otlploghttp.Exporter, error) {
	headers := internal.GetHeaders()

	exporter, err := otlploghttp.New(ctx,
		otlploghttp.WithEndpoint(constants.TelemetryEndpoint),
		otlploghttp.WithURLPath(constants.TelemetryLogsPath),
		otlploghttp.WithHeaders(headers),
	)
	if err != nil {
		return nil, err
	}

	return exporter, nil
}

func newProcessor(exporter *otlploghttp.Exporter) *log.BatchProcessor {
	return log.NewBatchProcessor(exporter)
}

func newLoggerProvider(processor log.Processor) *log.LoggerProvider {
	resource := internal.GetResource()
	loggerProvider := log.NewLoggerProvider(
		log.WithResource(resource),
		log.WithProcessor(processor),
	)

	return loggerProvider
}
