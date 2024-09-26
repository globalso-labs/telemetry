/*
 * telemetry
 * exporters.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 24 Sep 2024 01:36:05 -0500 by nick.
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

	"go.globalso.dev/x/telemetry/config"
)

func marshalExporters(ctx context.Context, telemetry *config.Telemetry) (map[string]interface{}, error) {
	exporters := make(map[string]interface{})

	exporters["debug"] = marshalDebugExporter(ctx, telemetry)

	switch telemetry.Protocol {
	// case "grpc":
	// 	return map[string]interface{}{"otlp": marshalGRPCExporter(ctx, telemetry)}, nil
	case "http", "https":
		exporters["otlphttp"] = marshalHTTPExporter(ctx, telemetry)
		exporters["prometheusremotewrite"] = marshalPrometheusRemoteWriteExporter(ctx, telemetry)
	default:
		return nil, fmt.Errorf("unsupported protocol: %s", telemetry.Protocol)
	}

	return exporters, nil
}

func marshalPrometheusRemoteWriteExporter(_ context.Context, telemetry *config.Telemetry) map[string]interface{} {
	var exporter = make(map[string]interface{})
	exporter["endpoint"] = fmt.Sprintf("%s://%s/%s", telemetry.Protocol, telemetry.Endpoint, "api/v1/push")
	exporter["headers"] = telemetry.Headers

	return exporter
}

func marshalDebugExporter(_ context.Context, _ *config.Telemetry) map[string]interface{} {
	var exporter = make(map[string]interface{})
	exporter["verbosity"] = "detailed"

	return exporter
}

func marshalHTTPExporter(_ context.Context, telemetry *config.Telemetry) map[string]interface{} {
	var exporter = make(map[string]interface{})
	exporter["logs_endpoint"] = fmt.Sprintf("%s://%s/%s", telemetry.Protocol, telemetry.Endpoint, telemetry.Logger.Path)
	exporter["traces_endpoint"] = fmt.Sprintf("%s://%s/%s", telemetry.Protocol, telemetry.Endpoint, telemetry.Tracer.Path)
	exporter["headers"] = telemetry.Headers

	exporter["sending_queue"] = map[string]interface{}{
		"storage": "file_storage",
	}

	return exporter
}
