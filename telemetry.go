/*
 * telemetry
 * telemetry.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Wed, 31 Jul 2024 20:24:08 -0500 by nick.
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

package telemetry

import (
	"context"
	"errors"

	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/logger"
	"go.globalso.dev/x/telemetry/meter"
	"go.globalso.dev/x/telemetry/shared"
	"go.globalso.dev/x/telemetry/tracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// Handle holds initialized telemetry components and provides shutdown helpers.
type Handle struct {
	Logger *logger.Logger
	Meter  *meter.Meter
	Tracer *tracer.Tracer
}

// Shutdown gracefully shuts down all initialized components, aggregating errors.
func (h *Handle) Shutdown(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	var errs []error
	if h == nil {
		return nil
	}
	if h.Meter != nil {
		if err := h.Meter.Shutdown(ctx); err != nil {
			errs = append(errs, err)
		}
	}
	if h.Tracer != nil {
		if err := h.Tracer.Shutdown(ctx); err != nil {
			errs = append(errs, err)
		}
	}
	if h.Logger != nil {
		if err := h.Logger.Shutdown(ctx); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

// Close is a convenience wrapper that shuts down with a background context.
func (h *Handle) Close() error {
	return h.Shutdown(context.Background())
}

// Start initializes telemetry and returns a handle to shut it down later.
// It is safe to call when telemetry or its subcomponents are disabled.
func Start(ctx context.Context, telemetry *config.Telemetry) (*Handle, error) {
	if telemetry == nil {
		telemetry = config.Default()
	}
	if telemetry.Resource == nil {
		telemetry.Resource = shared.NewResource()
	}
	if !telemetry.Enabled {
		return &Handle{}, nil
	}

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	handle := &Handle{}

	var err error
	if telemetry.Logger.Enabled {
		handle.Logger, err = logger.Initialize(ctx, telemetry)
		if err != nil {
			return nil, errors.Join(err, handle.Shutdown(ctx))
		}
	}

	if telemetry.Tracer.Enabled {
		handle.Tracer, err = tracer.Initialize(ctx, telemetry)
		if err != nil {
			return nil, errors.Join(err, handle.Shutdown(ctx))
		}
	}

	if telemetry.Meter.Enabled {
		handle.Meter, err = meter.Initialize(ctx, telemetry)
		if err != nil {
			return nil, errors.Join(err, handle.Shutdown(ctx))
		}
	}

	return handle, nil
}
