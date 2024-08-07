/*
 * telemetry
 * register.go
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

	"go.opentelemetry.io/contrib/instrumentation/host"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
)

var _handler = new(Holder)

// Register initializes and registers the OpenTelemetry meter provider.
//
// This function creates a new Holder instance using the provided context and configuration,
// sets it as the global meter provider, and starts the host and runtime instrumentation.
//
// Parameters:
// - ctx context.Context: The context for managing the meter's lifecycle.
// - opts *Options: The configuration for the meter.
//
// Returns:
// - error: An error if the meter initialization or instrumentation startup fails.
func Register(ctx context.Context, opts *Options) error {
	p, err := NewMeter(ctx, opts)
	if err != nil {
		return err
	}
	otel.SetMeterProvider(p.provider)

	if err = host.Start(); err != nil {
		return err
	}

	if err = runtime.Start(); err != nil {
		return err
	}

	_handler = p
	return nil
}

// Shutdown gracefully shuts down the OpenTelemetry meter provider.
//
// This function shuts down the global meter provider, ensuring that all pending
// telemetry data is flushed and resources are released.
//
// Parameters:
// - ctx context.Context: The context for managing the shutdown process.
//
// Returns:
// - error: An error if the shutdown process fails.
func Shutdown(ctx context.Context) error {
	return _handler.Shutdown(ctx)
}
