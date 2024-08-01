/*
 * telemetry
 * options_test.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Wed, 31 Jul 2024 22:11:29 -0500 by nick.
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

package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.globalso.dev/x/telemetry/common"
)

func Test_Options(t *testing.T) {
	t.Parallel()
	assert.NotNil(t, common.Options)

	var opts []common.Option
	opts = append(opts, common.WithID("test"))
	assert.NotEqual(t, "test", common.ID())

	opts = append(opts, common.WithName("name"))
	assert.NotEqual(t, "name", common.Name())

	opts = append(opts, common.WithVersion("version"))
	assert.NotEqual(t, "version", common.Version())

	opts = append(opts, common.WithOrganizationID("org"))
	assert.NotEqual(t, "org", common.OrganizationID())

	opts = append(opts, common.WithNamespace("namespace"))
	assert.NotEqual(t, "namespace", common.Namespace())

	opts = append(opts, common.WithOTLPEndpoint("endpoint"))
	assert.NotEqual(t, "endpoint", common.OTLPEndpoint())

	for _, opt := range opts {
		opt.Apply(&common.Options)
	}

	assert.Equal(t, "test", common.ID())
	assert.Equal(t, "name", common.Name())
	assert.Equal(t, "version", common.Version())
	assert.Equal(t, "namespace", common.Namespace())
	assert.Equal(t, "org", common.OrganizationID())
	assert.Equal(t, "endpoint", common.OTLPEndpoint())
}
