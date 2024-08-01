/*
 * telemetry
 * zerolog.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 9 Jul 2024 14:01:07 -0500 by nick.
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
	"context"

	"github.com/rs/zerolog"
	internal "go.globalso.dev/x/telemetry/logger/zerolog"
)

func init() { //nolint: gochecknoinits // This is the only way to set the default logger.
	zerolog.DefaultContextLogger = &internal.DefaultContextLogger
	zerolog.CallerMarshalFunc = internal.CallerMarshalFunc
}

// Ctx retrieves the zerolog.Logger from the provided context.
//
// This function extracts the zerolog.Logger instance associated with the given context.
// It is useful for logging within the context of a request or operation.
//
// Parameters:
// - ctx context.Context: The context from which to retrieve the logger.
//
// Returns:
// - *zerolog.Logger: The logger associated with the context.
func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}

// WithFields adds fields to the zerolog.Logger in the provided context.
//
// This function retrieves the zerolog.Logger from the context, adds the specified fields
// to the logger's context, and returns the updated logger. It is useful for adding
// contextual information to logs.
//
// Parameters:
// - ctx context.Context: The context from which to retrieve the logger.
// - fields map[string]any: The fields to add to the logger's context.
//
// Returns:
// - *zerolog.Logger: The updated logger with the added fields.
func WithFields(ctx context.Context, fields map[string]interface{}) *zerolog.Logger {
	l := Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Fields(fields)
	})

	return l
}

// Log logs a no-level message using the zerolog.Logger from the provided context.
func Log() *zerolog.Event {
	return internal.DefaultContextLogger.Log()
}

// Trace logs a trace level message using the zerolog.Logger from the provided context.
func Trace() *zerolog.Event { return internal.DefaultContextLogger.Trace() }

// Debug logs a debug level message using the zerolog.Logger from the provided context.
func Debug() *zerolog.Event { return internal.DefaultContextLogger.Debug() }

// Info logs an info level message using the zerolog.Logger from the provided context.
func Info() *zerolog.Event { return internal.DefaultContextLogger.Info() }

// Warn logs a warn level message using the zerolog.Logger from the provided context.
func Warn() *zerolog.Event { return internal.DefaultContextLogger.Warn() }

// Error logs an error level message using the zerolog.Logger from the provided context.
func Error() *zerolog.Event { return internal.DefaultContextLogger.Error() }

// Fatal logs a fatal level message using the zerolog.Logger from the provided context.
func Fatal() *zerolog.Event { return internal.DefaultContextLogger.Fatal() }

// Panic logs a panic level message using the zerolog.Logger from the provided context.
func Panic() *zerolog.Event { return internal.DefaultContextLogger.Panic() }
