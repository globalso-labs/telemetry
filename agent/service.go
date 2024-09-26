/*
 * telemetry
 * service.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 24 Sep 2024 03:30:57 -0500 by nick.
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

	"go.globalso.dev/x/telemetry/config"
)

// marshalService marshals the telemetry service configuration into a map.
//
// Parameters:
// - ctx context.Context: The context to use for marshaling.
// - telemetry *config.Telemetry: The telemetry configuration to be marshaled.
//
// Returns:
// - map[string]interface{}: The marshaled service configuration.
// - error: An error if the marshaling fails.
func marshalService(ctx context.Context, telemetry *config.Telemetry) map[string]interface{} {
	service := make(map[string]interface{})

	service["extensions"] = []string{"file_storage"}
	service["pipelines"] = marshalPipelines(ctx, telemetry)

	return service
}

// marshalPipelines marshals the telemetry pipelines configuration into a map.
//
// Parameters:
// - ctx context.Context: The context to use for marshaling.
// - telemetry *config.Telemetry: The telemetry configuration to be marshaled.
//
// Returns:
// - map[string]interface{}: The marshaled pipelines configuration.
func marshalPipelines(ctx context.Context, telemetry *config.Telemetry) map[string]interface{} {
	pipelines := make(map[string]interface{})

	pipelines["metrics"] = marshalMetricsPipeline(ctx, telemetry)
	pipelines["traces"] = marshalTracesPipeline(ctx, telemetry)
	pipelines["logs"] = marshalLogsPipeline(ctx, telemetry)

	return pipelines
}

// marshalLogsPipeline marshals the logs pipeline configuration into a map.
//
// Parameters:
// - ctx context.Context: The context to use for marshaling (unused).
// - telemetry *config.Telemetry: The telemetry configuration to be marshaled (unused).
//
// Returns:
// - map[string]interface{}: The marshaled logs pipeline configuration.
func marshalLogsPipeline(_ context.Context, _ *config.Telemetry) map[string]interface{} {
	logs := make(map[string]interface{})
	logs["exporters"] = []string{"otlphttp"}
	logs["receivers"] = []string{"filelog"}
	logs["processors"] = []string{"batch", "memory_limiter", "resourcedetection", "resource"}

	return logs
}

// marshalTracesPipeline marshals the traces pipeline configuration into a map.
//
// Parameters:
// - ctx context.Context: The context to use for marshaling (unused).
// - telemetry *config.Telemetry: The telemetry configuration to be marshaled (unused).
//
// Returns:
// - map[string]interface{}: The marshaled traces pipeline configuration.
func marshalTracesPipeline(_ context.Context, _ *config.Telemetry) map[string]interface{} {
	traces := make(map[string]interface{})
	traces["exporters"] = []string{"otlphttp"}
	traces["receivers"] = []string{"nop"}
	traces["processors"] = []string{"batch", "memory_limiter", "resourcedetection", "resource"}

	return traces
}

// marshalMetricsPipeline marshals the metrics pipeline configuration into a map.
//
// Parameters:
// - ctx context.Context: The context to use for marshaling (unused).
// - telemetry *config.Telemetry: The telemetry configuration to be marshaled (unused).
//
// Returns:
// - map[string]interface{}: The marshaled metrics pipeline configuration.
func marshalMetricsPipeline(_ context.Context, _ *config.Telemetry) map[string]interface{} {
	metrics := make(map[string]interface{})
	metrics["exporters"] = []string{"prometheusremotewrite"}
	metrics["receivers"] = []string{"hostmetrics"}
	metrics["processors"] = []string{"batch", "memory_limiter", "resourcedetection", "resource"}

	return metrics
}
