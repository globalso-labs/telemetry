/*
 * telemetry
 * meter_test.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 17 Sep 2024 23:01:10 -0500 by nick.
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
package meter_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/internal"
	"go.globalso.dev/x/telemetry/meter"
	"go.globalso.dev/x/telemetry/pkg/errors"
)

func TestInitialize_Success(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cfg := config.Default()

	m, err := meter.Initialize(ctx, cfg, internal.NewResource())
	defer m.Shutdown(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, m)

}

func TestInitialize_TelemetryNotEnabled(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cfg := &config.Telemetry{
		Enabled: false,
		Meter:   config.MeterDefault(),
	}

	m, err := meter.Initialize(ctx, cfg, internal.NewResource())
	assert.ErrorIs(t, err, errors.ErrTelemetryNotEnabled)
	assert.Nil(t, m)
}

func TestInitialize_MeterNotEnabled(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cfg := config.Default()
	cfg.Meter.Disable()

	m, err := meter.Initialize(ctx, cfg, internal.NewResource())
	assert.ErrorIs(t, err, errors.ErrTelemetryMeterNotEnabled)
	assert.Nil(t, m)
}
