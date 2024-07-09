/*
 * telemetry
 * options.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Mon, 8 Jul 2024 21:57:19 -0500 by nick.
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

package config

import (
	"go.globalso.dev/x/telemetry/common"
)

// Option interface with methods to apply options.
type Option interface {
	ApplyOption(*Config)
}

// genericOption struct with a function to apply options.
type genericOption struct {
	fn func(*common.Common)
}

// ApplyOption applies the genericOption to the given Config instance.
func (o *genericOption) ApplyOption(cfg *Config) {
	o.fn(cfg.Meter.Common())
	o.fn(cfg.Common())
}

// NewGenericOption creates a new Option instance.
func newGenericOption(fn func(*common.Common)) Option {
	return &genericOption{fn: fn}
}

func WithOrganizationID(id string) Option {
	return newGenericOption(func(t *common.Common) {
		t.OrganizationID = id
	})
}

func WithOTLPEndpoint(endpoint string) Option {
	return newGenericOption(func(t *common.Common) {
		t.OTLPEndpoint = endpoint
	})
}
