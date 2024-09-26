/*
 * telemetry
 * agent.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Thu, 19 Sep 2024 23:56:43 -0500 by nick.
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

package agent

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/logger"
	"go.globalso.dev/x/telemetry/pkg/errors"
	"go.mau.fi/zerozap"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/provider/yamlprovider"
	"go.opentelemetry.io/collector/otelcol"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Initialize(ctx context.Context, telemetry config.Telemetry) (*Agent, error) {
	info := component.BuildInfo{
		Command:     "agent",
		Description: fmt.Sprint("OpenTelemetry Collector Agent on ", telemetry.Resource.GetName()),
		Version:     telemetry.Resource.GetVersion(),
	}

	if !telemetry.Enabled {
		return nil, errors.ErrTelemetryNotEnabled
	}

	encoded, err := marshalTelemetryConfig(ctx, &telemetry)
	if err != nil {
		return nil, err
	}

	settings := otelcol.CollectorSettings{
		Factories:               components,
		BuildInfo:               info,
		DisableGracefulShutdown: false,
		ConfigProviderSettings: otelcol.ConfigProviderSettings{
			ResolverSettings: confmap.ResolverSettings{ //nolint:exhaustruct // We don't need all fields.
				URIs: []string{"yaml:" + string(encoded)},
				ProviderFactories: []confmap.ProviderFactory{
					yamlprovider.NewFactory(),
				},
			},
		},
		LoggingOptions: []zap.Option{
			zap.WrapCore(func(_ zapcore.Core) zapcore.Core {
				return zerozap.New(
					logger.Ctx(ctx).With().
						CallerWithSkipFrameCount(5). //nolint:mnd // We need to 5 frames to get the caller.
						Logger().Level(zerolog.InfoLevel),
				)
			}),
		},
		SkipSettingGRPCLogger: true,
	}

	collector, err := otelcol.NewCollector(settings)
	if err != nil {
		return nil, err
	}

	return &Agent{
		settings:  settings,
		collector: collector,
	}, nil
}
