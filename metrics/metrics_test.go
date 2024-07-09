/*
 * telemetry
 * metrics_test.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 9 Jul 2024 01:26:40 -0500 by nick.
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

package metrics_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.globalso.dev/x/telemetry/metrics"
)

func Test_RegisterGlobal(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cfg := metrics.NewConfig()

	metrics.Register(ctx, &cfg)
	defer metrics.Shutdown(ctx)
}

func Test_NewMeter(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cfg := metrics.NewConfig()

	m, err := metrics.NewMeter(ctx, &cfg)
	defer m.Shutdown(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, m)

	err = m.Provider().ForceFlush(ctx)
	assert.Nil(t, err)
}
