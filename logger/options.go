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

package logger

import (
	"io"
	"time"

	"github.com/rs/zerolog"
	"go.globalso.dev/x/telemetry/common"
	"go.globalso.dev/x/telemetry/logger/constants"
)

// Options holds configuration options for a meter.
type Options struct {
	Enabled bool // Enabled specifies whether the meter is enabled.

	ExportInterval time.Duration // ExportInterval specifies the interval at which the meter exports data.
	Level          zerolog.Level // Level specifies the log constants of the meter.

	Writer io.Writer // Writer specifies the writer to write logs to.
}

// Option defines an interface for applying meter options.
type Option interface {
	ApplyLoggerOption(*Options)
}

// loggerOption implements the Option interface with a function to apply a configuration option.
type loggerOption struct {
	fn func(*Options) // fn is a function that applies a configuration option to Options.
}

// ApplyLoggerOption applies a configuration option to the given Options instance.
func (o *loggerOption) ApplyLoggerOption(option *Options) {
	o.fn(option)
}

// newLoggerOption creates a new loggerOption with the specified function.
// The function provided should define how to modify the Options instance.
func newLoggerOption(fn func(*Options)) Option { //nolint:ireturn
	return &loggerOption{fn: fn}
}

// WithLevel returns an Option that sets the log constants of the meter.
// This function is a convenient way to specify the log level of the meter.
func WithLevel(l zerolog.Level) Option { //nolint:ireturn
	return newLoggerOption(func(o *Options) {
		o.Level = l
	})
}

// WithVerbosity returns an Option that sets the log constants of the meter based on verbosity.
// This function is a convenient way to specify the log level of the meter based on verbosity.
func WithVerbosity(verbosity int) Option { //nolint:ireturn
	return newLoggerOption(func(o *Options) {
		vLevel := constants.DefaultLoggerLevel - zerolog.Level(verbosity)
		if verbosity < int(zerolog.TraceLevel) {
			vLevel = zerolog.TraceLevel
		}
		o.Level = vLevel
	})
}

// WithWriter returns an Option that sets the writer for the meter.
// This function is a convenient way to specify the writer to which logs are written.
func WithWriter(w io.Writer) Option { //nolint:ireturn // Used to create an Option.
	return newLoggerOption(func(o *Options) {
		o.Writer = w
	})
}

// WithEnabled returns an Option that sets the enabled state of the meter.
// This function is a convenient way to enable or disable the meter functionality.
func WithEnabled(enabled bool) Option { //nolint:ireturn
	return newLoggerOption(func(o *Options) {
		o.Enabled = enabled
	})
}

// WithExportInterval returns an Option that sets the read interval of the meter.
// This function is a convenient way to specify how often the meter reads data.
func WithExportInterval(interval time.Duration) Option { //nolint:ireturn
	return newLoggerOption(func(o *Options) {
		o.ExportInterval = interval
	})
}

// WithCommonOptions returns an Option that applies common options to the meter.
// This function is a convenient way to apply common options to the meter.
func WithCommonOptions(opts ...common.Option) Option { //nolint:ireturn
	return newLoggerOption(func(_ *Options) {
		for _, opt := range opts {
			opt.Apply(&common.Options)
		}
	})
}
