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

	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/internal"
	"go.globalso.dev/x/telemetry/logger"
	"go.globalso.dev/x/telemetry/meter"
	"go.globalso.dev/x/telemetry/tracer"
)

// Telemetry represents the telemetry system, including logger, tracer, and meter.
type Telemetry struct {
	ctx    context.Context
	config *config.Telemetry

	logger *logger.Logger
	tracer *tracer.Tracer
	meter  *meter.Meter
}

// GetConfig returns the telemetry configuration.
//
// Returns:
// - *config.Telemetry: The telemetry configuration.
func (t *Telemetry) GetConfig() *config.Telemetry {
	return t.config
}

// GetResource returns the telemetry resource configuration.
//
// Returns:
// - *internal.Resource: The telemetry resource configuration.
func (t *Telemetry) GetResource() *internal.Resource {
	return t.config.Resource
}

// newTelemetry creates a new Telemetry instance with default configuration.
//
// Parameters:
// - ctx context.Context: The context to use for the telemetry instance.
//
// Returns:
// - *Telemetry: The new Telemetry instance.
func newTelemetry(ctx context.Context) *Telemetry {
	return &Telemetry{
		ctx:    ctx,
		config: config.Default(),
		logger: nil,
		tracer: nil,
		meter:  nil,
	}
}

// Initialize sets up the Telemetry instance with the provided options.
//
// Parameters:
// - ctx context.Context: The context to use for initialization.
// - opts ...Option: The options to configure the telemetry instance.
//
// Returns:
// - *Telemetry: The initialized Telemetry instance.
// - error: An error if the initialization fails.
func Initialize(ctx context.Context, opts ...Option) (*Telemetry, error) {
	t := newTelemetry(ctx)
	for _, opt := range opts {
		opt(t)
	}

	var err error
	t.logger, err = logger.Initialize(t.ctx, t.config)
	if err != nil {
		return nil, err
	}

	t.tracer, err = tracer.Initialize(t.ctx, t.config)
	if err != nil {
		return nil, err
	}

	t.meter, err = meter.Initialize(t.ctx, t.config)
	if err != nil {
		return nil, err
	}

	return t, nil
}
