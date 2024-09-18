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

	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/internal"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Tracer struct {
	provider  *trace.TracerProvider
	processor trace.SpanProcessor
	exporter  trace.SpanExporter
}

func (t Tracer) Provider() *trace.TracerProvider {
	return t.provider
}

func (t Tracer) Shutdown(ctx context.Context) error {
	if err := t.processor.Shutdown(ctx); err != nil {
		return err
	}

	if err := t.exporter.Shutdown(ctx); err != nil {
		return err
	}

	if err := t.provider.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func newExporter(ctx context.Context, cfg *config.Telemetry) (*otlptrace.Exporter, error) {
	return otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(cfg.Endpoint),
		otlptracehttp.WithURLPath(cfg.Tracer.Path),
		otlptracehttp.WithHeaders(cfg.Headers),
	)
}

func newProcessor(_ context.Context, exporter *otlptrace.Exporter) trace.SpanProcessor { //nolint:ireturn
	return trace.NewBatchSpanProcessor(exporter)
}

func newProvider(ctx context.Context, res *internal.Resource, processor trace.SpanProcessor) *trace.TracerProvider {
	resource := internal.GetResource(ctx, res)
	return trace.NewTracerProvider(
		trace.WithResource(resource),
		trace.WithSpanProcessor(processor),
	)
}
