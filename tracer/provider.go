/*
 * telemetry
 * provider.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Sun, 4 Aug 2024 02:21:28 -0500 by nick.
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

package tracer

import (
	"context"
	"fmt"

	"go.globalso.dev/x/telemetry/internal"
	"go.globalso.dev/x/telemetry/internal/constants"
	"go.globalso.dev/x/telemetry/pkg/errors"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Holder struct {
	provider  *trace.TracerProvider
	processor trace.SpanProcessor
	exporter  trace.SpanExporter
}

func NewTracer(ctx context.Context, config *Options) (*Holder, error) {
	if !config.IsEnabled() {
		return nil, errors.ErrTelemetryTracesNotEnabled
	}

	holder := new(Holder)

	// Create the exporter.
	exporter, err := newHTTPExporter(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}
	holder.exporter = exporter

	// Create the processor.
	holder.processor = newProcessor(exporter)

	// Create the tracer provider.
	holder.provider = newProvider(holder.processor)

	return holder, nil
}

func (t Holder) Provider() *trace.TracerProvider {
	return t.provider
}

func (t Holder) Shutdown(ctx context.Context) error {
	if err := t.provider.Shutdown(ctx); err != nil {
		return err
	}

	if err := t.processor.Shutdown(ctx); err != nil {
		return err
	}

	if err := t.exporter.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func newHTTPExporter(ctx context.Context, _ *Options) (*otlptrace.Exporter, error) {
	return otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(constants.TelemetryEndpoint),
		otlptracehttp.WithURLPath(constants.TelemetryTracesPath),
		otlptracehttp.WithHeaders(internal.GetHeaders()),
	)
}

func newProcessor(exporter *otlptrace.Exporter) trace.SpanProcessor { //nolint:ireturn
	return trace.NewBatchSpanProcessor(exporter)
}

func newProvider(processor trace.SpanProcessor) *trace.TracerProvider {
	resource := internal.GetResource()
	return trace.NewTracerProvider(
		trace.WithResource(resource),
		trace.WithSpanProcessor(processor),
	)
}
