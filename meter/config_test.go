/*
 * telemetry
 * config_test.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Mon, 8 Jul 2024 21:07:18 -0500 by nick.
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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.globalso.dev/x/telemetry/meter"
)

func Test_NewConfig(t *testing.T) {
	t.Parallel()

	cfg := meter.NewConfig()
	assert.NotNil(t, cfg)
	assert.True(t, cfg.IsEnabled())

}

func Test_NewConfigWithInternalOptions(t *testing.T) {
	t.Parallel()

	t.Run("enabled", func(t *testing.T) {
		cfg := meter.NewConfig(meter.WithEnabled(true))
		assert.NotNil(t, cfg)
		assert.True(t, cfg.IsEnabled())
	})

	t.Run("disabled", func(t *testing.T) {
		cfg := meter.NewConfig(meter.WithEnabled(false))
		assert.NotNil(t, cfg)
		assert.False(t, cfg.IsEnabled())
	})

	t.Run("read_interval", func(t *testing.T) {
		cfg := meter.NewConfig(meter.WithReadInterval(10 * time.Second))
		assert.NotNil(t, cfg)
		assert.True(t, cfg.IsEnabled())
		assert.Equal(t, 10*time.Second, cfg.ReadInterval)
	})
}