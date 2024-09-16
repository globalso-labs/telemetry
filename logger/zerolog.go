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
)

// Ctx retrieves the zerolog.Logger from the provided context.
func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}

// With returns a new zerolog.Context.
func With() zerolog.Context { return logger.With() }

// Log logs a no-constants message using the zerolog.Logger from the provided context.
func Log() *zerolog.Event {
	return logger.Log()
}

// Trace logs a trace constants message using the zerolog.Logger from the provided context.
func Trace() *zerolog.Event { return logger.Trace() }

// Debug logs a debug constants message using the zerolog.Logger from the provided context.
func Debug() *zerolog.Event { return logger.Debug() }

// Info logs an info constants message using the zerolog.Logger from the provided context.
func Info() *zerolog.Event { return logger.Info() }

// Warn logs a warn constants message using the zerolog.Logger from the provided context.
func Warn() *zerolog.Event { return logger.Warn() }

// Error logs an error constants message using the zerolog.Logger from the provided context.
func Error() *zerolog.Event { return logger.Error() }

// Fatal logs a fatal constants message using the zerolog.Logger from the provided context.
func Fatal() *zerolog.Event { return logger.Fatal() }

// Panic logs a panic constants message using the zerolog.Logger from the provided context.
func Panic() *zerolog.Event { return logger.Panic() }
