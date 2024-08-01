/*
 * telemetry
 * register.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 9 Jul 2024 01:45:28 -0500 by nick.
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

	"go.globalso.dev/x/telemetry/logger/otlp"
	internal "go.globalso.dev/x/telemetry/logger/zerolog"
	"go.opentelemetry.io/otel/log/global"
)

var _handler = new(Logger)

func Register(ctx context.Context, opts *Options) error {
	p, err := NewLogger(ctx, opts)
	if err != nil {
		return err
	}

	global.SetLoggerProvider(p.provider)

	writer := internal.WithCloser(opts.Writer, p.Close)
	l := internal.DefaultContextLogger.
		Output(writer).
		Hook(otlp.Hook{}).
		Level(internal.FromLevel(opts.Level))
	internal.DefaultContextLogger = l

	_handler = p
	return nil
}

func Shutdown(ctx context.Context) error {
	return _handler.Shutdown(ctx)
}
