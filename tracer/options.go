/*
 * telemetry
 * options_generic.go
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

package meter

import (
	"time"

	"go.globalso.dev/x/telemetry/common"
)

// Options holds configuration options for a meter.
type Options struct {
	Enabled        bool          // Enabled specifies whether the meter is enabled.
	ReadInterval   time.Duration // ReadInterval specifies the interval at which the meter reads data.
	ExportInterval time.Duration // ExportInterval specifies the interval at which the meter exports data.
}

// Option defines an interface for applying meter options.
type Option interface {
	ApplyOption(*Options)
}

// meterOption implements the Option interface with a function to apply a configuration option.
type meterOption struct {
	fn func(*Options) // fn is a function that applies a configuration option to Options.
}

// ApplyOption applies a configuration option to the given Options instance.
func (o *meterOption) ApplyOption(option *Options) {
	o.fn(option)
}

func newMeterOption(fn func(*Options)) Option { //nolint:ireturn
	return &meterOption{fn: fn}
}

func WithCommonOptions(opts ...common.Option) Option { //nolint:ireturn
	return newMeterOption(func(_ *Options) {
		for _, opt := range opts {
			opt.Apply(&common.Options)
		}
	})
}

func WithEnabled(enabled bool) Option { //nolint:ireturn
	return newMeterOption(func(o *Options) {
		o.Enabled = enabled
	})
}

func WithReadInterval(interval time.Duration) Option { //nolint:ireturn
	return newMeterOption(func(o *Options) {
		o.ReadInterval = interval
	})
}
