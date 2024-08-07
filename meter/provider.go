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
	"fmt"

	"go.globalso.dev/x/telemetry/internal"
	"go.globalso.dev/x/telemetry/internal/constants"
	"go.globalso.dev/x/telemetry/pkg/errors"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/sdk/metric"
)

// Holder represents a structure for managing meter collection and export.
// It contains the following fields:
// - provider: A pointer to the MeterProvider which manages metric instruments and readers.
// - reader: A metric.Reader that periodically reads and exports meter.
// - exporter: A metric.Exporter that sends meter to a backend.
type Holder struct {
	provider *metric.MeterProvider
	reader   metric.Reader
	exporter metric.Exporter
}

// NewMeter creates a new Holder instance for meter collection and export.
// It takes a context and a Holder configuration as parameters.
// If telemetry is not enabled in the configuration, it returns an error.
// It initializes the metric exporter, reader, and provider, and returns
// the Holder instance or an error if any step fails.
func NewMeter(ctx context.Context, config *Options) (*Holder, error) {
	if !config.IsEnabled() {
		return nil, errors.ErrTelemetryMetricsNotEnabled
	}

	holder := new(Holder)

	// Create the exporter.
	exporter, err := newHTTPExporter(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %w", err)
	}
	holder.exporter = exporter

	// Create the reader.
	holder.reader = newReader(exporter, config)

	// Create the provider.
	holder.provider = newProvider(holder.reader)

	return holder, nil
}

func (m Holder) Provider() *metric.MeterProvider {
	return m.provider
}

// Shutdown stops the metric provider.
func (m Holder) Shutdown(ctx context.Context) error {
	if err := m.provider.Shutdown(ctx); err != nil {
		return err
	}

	if err := m.exporter.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

// newHTTPExporter creates a new OTLP HTTP exporter for meter.
// It takes a context and a Holder configuration as parameters.
// It returns an OTLP HTTP exporter or an error if the creation fails.
func newHTTPExporter(ctx context.Context, _ *Options) (*otlpmetrichttp.Exporter, error) {
	return otlpmetrichttp.New(ctx,
		otlpmetrichttp.WithEndpoint(constants.TelemetryEndpoint),
		otlpmetrichttp.WithURLPath(constants.TelemetryMetricsPath),
		otlpmetrichttp.WithHeaders(internal.GetHeaders()),
	)
}

// newReader creates a new PeriodicReader for meter.
// It takes a metric exporter and a Holder configuration as parameters.
// It returns a PeriodicReader configured with the specified export interval.
func newReader(exporter metric.Exporter, opts *Options) *metric.PeriodicReader {
	return metric.NewPeriodicReader(exporter,
		metric.WithInterval(opts.ExportInterval),
	)
}

// newProvider creates a new MeterProvider for meter.
// It takes a metric reader and a Holder configuration as parameters.
// It returns a MeterProvider configured with the specified resource and reader.
func newProvider(reader metric.Reader) *metric.MeterProvider {
	resource := internal.GetResource()

	return metric.NewMeterProvider(
		metric.WithResource(resource),
		metric.WithReader(reader),
	)
}
