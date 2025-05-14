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

	"go.globalso.dev/x/telemetry/common"
	"go.globalso.dev/x/telemetry/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/sdk/log"
)

type Logger struct {
	provider  *log.LoggerProvider
	processor log.Processor
	exporter  log.Exporter
}

func (l *Logger) Provider() *log.LoggerProvider {
	return l.provider
}

func (l *Logger) Shutdown(ctx context.Context) error {
	if err := l.processor.Shutdown(ctx); err != nil {
		return err
	}
	if err := l.exporter.Shutdown(ctx); err != nil {
		return err
	}
	if err := l.provider.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (l *Logger) Close() error {
	ctx := context.Background()
	return l.Shutdown(ctx)
}

func newExporter(ctx context.Context, cfg *config.Telemetry) (*otlploghttp.Exporter, error) {
	return otlploghttp.New(ctx,
		otlploghttp.WithEndpoint(cfg.Endpoint),
		otlploghttp.WithURLPath(cfg.Logger.Path),
		otlploghttp.WithHeaders(cfg.Headers),
	)
}

func newProcessor(_ context.Context, exporter *otlploghttp.Exporter) *log.BatchProcessor {
	return log.NewBatchProcessor(exporter)
}

func newLoggerProvider(ctx context.Context, res *common.Resource, processor log.Processor) *log.LoggerProvider {
	resource := common.GetResource(ctx, res)

	return log.NewLoggerProvider(
		log.WithResource(resource),
		log.WithProcessor(processor),
	)
}
