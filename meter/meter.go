/*
 * telemetry
 * meter.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 17 Sep 2024 22:46:10 -0500 by nick.
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

	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
)

// Initialize sets up the Meter instance with the provided context, configuration, and resource.
//
// Parameters:
// - ctx context.Context: The context to use for initialization.
// - cfg *config.Telemetry: The telemetry configuration to be used.
// - res *internal.Resource: The resource to be used.
//
// Returns:
// - *Meter: The initialized Meter instance.
// - error: An error if the initialization fails.
func Initialize(ctx context.Context, telemetry *config.Telemetry) (*Meter, error) {
	if !telemetry.Enabled {
		return nil, errors.ErrTelemetryNotEnabled
	}

	if !telemetry.Meter.Enabled {
		return nil, errors.ErrTelemetryMeterNotEnabled
	}

	holder := new(Meter)

	// Create the exporter.
	exporter, err := newExporter(ctx, telemetry)
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %w", err)
	}
	holder.exporter = exporter

	// Create the reader.
	holder.reader = newReader(ctx, exporter, telemetry)

	// Create the provider.
	holder.provider = newProvider(ctx, telemetry.Resource, holder.reader)

	// Set the global meter.
	otel.SetMeterProvider(holder.provider)

	// Run runtime collectors
	_ = runtime.Start()

	return holder, nil
}
