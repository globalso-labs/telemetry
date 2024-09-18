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

type Telemetry struct {
	ctx      context.Context
	config   *config.Telemetry
	resource *internal.Resource

	logger *logger.Logger
	tracer *tracer.Tracer
	meter  *meter.Meter
}

func (t *Telemetry) GetConfig() *config.Telemetry {
	return t.config
}

func (t *Telemetry) GetResource() *internal.Resource {
	return t.resource
}

func newTelemetry(ctx context.Context) *Telemetry {
	return &Telemetry{
		ctx:      ctx,
		config:   config.Default(),
		resource: nil,
	}
}

func Initialize(ctx context.Context, opts ...Option) *Telemetry {
	t := newTelemetry(ctx)
	for _, opt := range opts {
		opt(t)
	}

	return t
}
