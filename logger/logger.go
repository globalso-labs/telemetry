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
	"fmt"

	"github.com/rs/zerolog"
	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/logger/drivers"
	"go.globalso.dev/x/telemetry/logger/hooks"
	custom "go.globalso.dev/x/telemetry/logger/zerolog"
	"go.globalso.dev/x/telemetry/pkg/errors"
	"go.opentelemetry.io/otel/log/global"
)

type Instance = zerolog.Logger

func Initialize(ctx context.Context, telemetry *config.Telemetry) (*Logger, error) {
	if !telemetry.Enabled {
		return nil, errors.ErrTelemetryNotEnabled
	}

	if !telemetry.Logger.Enabled {
		return nil, errors.ErrTelemetryLoggerNotEnabled
	}

	holder := new(Logger)

	// Create the exporter.
	exporter, err := newExporter(ctx, telemetry)
	if err != nil {
		return nil, fmt.Errorf("failed to create log exporter: %w", err)
	}
	holder.exporter = exporter

	// Create the processor.
	holder.processor = newProcessor(ctx, exporter)

	// Create the holder provider.
	holder.provider = newLoggerProvider(ctx, telemetry.Resource, holder.processor)

	// Set the global logger.
	global.SetLoggerProvider(holder.provider)

	// Get the writer driver.
	driver, err := drivers.New(telemetry.Logger.Drivers)
	if err != nil {
		return nil, fmt.Errorf("failed to create log driver: %w", err)
	}

	// Wrap the driver with the closer.
	writer := WithCloser(driver, holder.Close)

	// Parse level
	level, err := zerolog.ParseLevel(telemetry.Logger.Level)
	if err != nil {
		return nil, fmt.Errorf("failed to parse log level: %w", err)
	}

	hook, err := hooks.New(telemetry.Logger.Hooks)
	if err != nil {
		return nil, fmt.Errorf("failed to create log hook: %w", err)
	}

	// Create the logger.
	logger = zerolog.New(writer).Level(level).Hook(hook...).With().Timestamp().Caller().Logger()

	// Set the caller marshal function.
	zerolog.DefaultContextLogger = &logger
	zerolog.CallerMarshalFunc = custom.CallerMarshalFunc

	return holder, nil
}
