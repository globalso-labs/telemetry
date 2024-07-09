/*
 * telemetry
 * config_test.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Mon, 8 Jul 2024 20:42:52 -0500 by nick.
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

package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/internal"
)

func Test_New(t *testing.T) {
	t.Parallel()

	c := config.New()
	assert.NotNil(t, c)
	assert.True(t, c.IsEnabled())
	assert.Equal(t, internal.TelemetryEndpoint, c.Common().OTLPEndpoint)
	assert.Equal(t, internal.DefaultOrganizationID, c.Common().OrganizationID)

}

func Test_NewWithOrganizationID(t *testing.T) {
	t.Parallel()

	c := config.New(config.WithOrganizationID("test"))
	assert.NotNil(t, c)

	assert.Equal(t, "test", c.Common().OrganizationID)
	assert.Equal(t, "test", c.Metrics.Common().OrganizationID)

	assert.Equal(t, internal.TelemetryEndpoint, c.Common().OTLPEndpoint)
	assert.Equal(t, internal.TelemetryEndpoint, c.Metrics.Common().OTLPEndpoint)
}

func Test_NewWithOTLPEndpoint(t *testing.T) {
	t.Parallel()

	c := config.New(config.WithOTLPEndpoint("http://localhost:4317"))
	assert.NotNil(t, c)

	assert.Equal(t, internal.DefaultOrganizationID, c.Common().OrganizationID)
	assert.Equal(t, internal.DefaultOrganizationID, c.Metrics.Common().OrganizationID)

	assert.Equal(t, "http://localhost:4317", c.Common().OTLPEndpoint)
	assert.Equal(t, "http://localhost:4317", c.Metrics.Common().OTLPEndpoint)
}
