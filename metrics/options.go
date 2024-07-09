/*
 * telemetry
 * options.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Mon, 8 Jul 2024 20:53:15 -0500 by nick.
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

package metrics

import (
	"time"

	"go.globalso.dev/x/telemetry/common"
	"go.globalso.dev/x/telemetry/internal/constants"
)

var defaultOptions = MeterOptions{
	common: common.DefaultOptions(),

	Enabled:        true,
	ReadInterval:   constants.DefaultMetricReadInterval,
	ExportInterval: constants.DefaultMetricExportInterval,
}

// MeterOptions holds configuration options for a meter.
type MeterOptions struct {
	common common.Common

	Enabled        bool          // Enabled specifies whether the meter is enabled.
	ReadInterval   time.Duration // ReadInterval specifies the interval at which the meter reads data.
	ExportInterval time.Duration // ExportInterval specifies the interval at which the meter exports data.
}

// MeterOption defines an interface for applying meter options.
type MeterOption interface {
	ApplyMeterOption(*MeterOptions)
}

// meterOption implements the MeterOption interface with a function to apply a configuration option.
type meterOption struct {
	fn func(*MeterOptions) // fn is a function that applies a configuration option to MeterOptions.
}

// ApplyMeterOption applies a configuration option to the given MeterOptions instance.
func (o *meterOption) ApplyMeterOption(option *MeterOptions) {
	o.fn(option)
}

// newMeterOption creates a new meterOption with the specified function.
// The function provided should define how to modify the MeterOptions instance.
func newMeterOption(fn func(*MeterOptions)) MeterOption { //nolint:ireturn // Used to create a MeterOption.
	return &meterOption{fn: fn}
}

// WithCommonOptions returns a MeterOption that sets the configuration options of the meter.
// This function is a convenient way to set the configuration options of the meter.
func WithCommonOptions(opts common.Common) MeterOption { //nolint:ireturn // Used to create a MeterOption.
	return newMeterOption(func(o *MeterOptions) {
		o.common = opts
	})
}

// WithEnabled returns a MeterOption that sets the enabled state of the meter.
// This function is a convenient way to enable or disable the meter functionality.
func WithEnabled(enabled bool) MeterOption { //nolint:ireturn // Used to create a MeterOption.
	return newMeterOption(func(o *MeterOptions) {
		o.Enabled = enabled
	})
}

// WithReadInterval returns a MeterOption that sets the read interval of the meter.
// This function is a convenient way to specify how often the meter reads data.
func WithReadInterval(interval time.Duration) MeterOption { //nolint:ireturn // Used to create a MeterOption.
	return newMeterOption(func(o *MeterOptions) {
		o.ReadInterval = interval
	})
}
