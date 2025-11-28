/*
 * telemetry
 * provider.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Mon, 8 Jul 2024 20:45:52 -0500 by nick.
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

package meter

import (
	"context"

	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/shared"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/sdk/metric"
)

// Meter represents a structure for managing meter collection and export.
// It contains the following fields:
// - provider: A pointer to the MeterProvider which manages metric instruments and readers.
// - reader: A metric.Reader that periodically reads and exports meter.
// - exporter: A metric.Exporter that sends meter to a backend.
type Meter struct {
	provider *metric.MeterProvider
	reader   metric.Reader
	exporter metric.Exporter
}

// Provider returns the MeterProvider associated with the Meter.
// It provides access to the metric instruments and readers managed by the provider.
func (m *Meter) Provider() *metric.MeterProvider {
	return m.provider
}

// Shutdown gracefully shuts down the Meter, including its provider, reader, and exporter.
// It takes a context to allow for cancellation and returns an error if any component fails to shut down.
func (m *Meter) Shutdown(ctx context.Context) error {
	if err := m.reader.Shutdown(ctx); err != nil {
		return err
	}
	if err := m.exporter.Shutdown(ctx); err != nil {
		return err
	}
	if err := m.provider.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

// newExporter creates a new OTLP HTTP exporter for meter.
// It takes a context and a Meter configuration as parameters.
// It returns an OTLP HTTP exporter or an error if the creation fails.
func newExporter(ctx context.Context, cfg *config.Telemetry) (*otlpmetrichttp.Exporter, error) {
	return otlpmetrichttp.New(ctx,
		otlpmetrichttp.WithEndpoint(cfg.Endpoint),
		otlpmetrichttp.WithURLPath(cfg.Meter.Path),
		otlpmetrichttp.WithHeaders(cfg.Headers),
	)
}

// newReader creates a new PeriodicReader for meter.
// It takes a metric exporter and a Meter configuration as parameters.
// It returns a PeriodicReader configured with the specified export interval.
func newReader(_ context.Context, exporter metric.Exporter, cfg *config.Telemetry) *metric.PeriodicReader {
	return metric.NewPeriodicReader(exporter,
		metric.WithInterval(cfg.Meter.Scrape.Interval),
	)
}

// newProvider creates a new MeterProvider for meter.
// It takes a metric reader and a Meter configuration as parameters.
// It returns a MeterProvider configured with the specified resource and reader.
func newProvider(ctx context.Context, res *shared.Resource, reader metric.Reader) *metric.MeterProvider {
	resource := shared.GetResource(ctx, res)
	return metric.NewMeterProvider(
		metric.WithResource(resource),
		metric.WithReader(reader),
	)
}
