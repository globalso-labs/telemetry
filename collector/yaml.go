/*
 * telemetry
 * yaml.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 24 Sep 2024 01:26:12 -0500 by nick.
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

package collector

import (
	"context"

	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/internal"
)

func marshalTelemetryConfig(ctx context.Context, telemetry *config.Telemetry) ([]byte, error) {
	extensions, err := marshalExtensions(ctx, telemetry)
	if err != nil {
		return nil, err
	}
	internal.Merge(telemetry.Agent.Extensions, extensions)

	exporters, err := marshalExporters(ctx, telemetry)
	if err != nil {
		return nil, err
	}
	internal.Merge(telemetry.Agent.Exporters, exporters)

	processors := marshalProcessors(ctx, telemetry)
	internal.Merge(telemetry.Agent.Processors, processors)

	receivers := marshalReceivers(ctx, telemetry)
	internal.Merge(telemetry.Agent.Receivers, receivers)

	service := marshalService(ctx, telemetry)
	internal.Merge(telemetry.Agent.Service, service)

	data, err := telemetry.Agent.Dump()
	if err != nil {
		return nil, err
	}

	return data, nil
}
