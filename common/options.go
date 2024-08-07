/*
 * telemetry
 * options_generic.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Mon, 8 Jul 2024 21:03:37 -0500 by nick.
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

package common

import (
	"go.globalso.dev/x/telemetry/internal/constants"
)

var defaultOptions = Common{
	id:             constants.MachineID,
	name:           constants.ServiceName,
	namespace:      constants.ServiceNamespace,
	version:        constants.ServiceVersion,
	organizationID: constants.DefaultOrganizationID,
	otlpEndpoint:   constants.TelemetryEndpoint,
}

var Options = defaultOptions

// Option defines an interface for applying meter options.
type Option interface {
	Apply(common *Common)
}

type option struct {
	fn func(*Common)
}

func (o *option) Apply(common *Common) {
	o.fn(common)
}

func newOption(fn func(*Common)) Option { //nolint:ireturn
	return &option{fn: fn}
}

func WithID(id string) Option { //nolint:ireturn
	return newOption(func(o *Common) {
		o.id = id
	})
}

func WithName(name string) Option { //nolint:ireturn
	return newOption(func(o *Common) {
		o.name = name
	})
}

func WithNamespace(namespace string) Option { //nolint:ireturn
	return newOption(func(o *Common) {
		o.namespace = namespace
	})
}

func WithVersion(version string) Option { //nolint:ireturn
	return newOption(func(o *Common) {
		o.version = version
	})
}

func WithOrganizationID(organizationID string) Option { //nolint:ireturn
	return newOption(func(o *Common) {
		o.organizationID = organizationID
	})
}

func WithOTLPEndpoint(endpoint string) Option { //nolint:ireturn
	return newOption(func(o *Common) {
		o.otlpEndpoint = endpoint
	})
}
