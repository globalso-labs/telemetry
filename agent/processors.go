/*
 * telemetry
 * processors.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 24 Sep 2024 03:43:56 -0500 by nick.
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

func marshalProcessors(ctx context.Context, telemetry *config.Telemetry) map[string]interface{} {
	processors := make(map[string]interface{})

	processors["batch"] = marshalBatchProcessor(ctx, telemetry)
	processors["memory_limiter"] = marshalMemoryLimiterProcessor(ctx, telemetry)
	processors["resourcedetection"] = marshalResourceDetectionProcessor(ctx, telemetry)
	processors["resource"] = marshalResourceProcessor(ctx, telemetry)

	return processors
}

func marshalMemoryLimiterProcessor(_ context.Context, _ *config.Telemetry) map[string]interface{} {
	processor := make(map[string]interface{})

	processor["check_interval"] = "10s"
	processor["limit_mib"] = 512

	return processor
}

func marshalBatchProcessor(_ context.Context, _ *config.Telemetry) map[string]interface{} {
	return nil
}

func marshalResourceDetectionProcessor(_ context.Context, _ *config.Telemetry) map[string]interface{} {
	processor := make(map[string]interface{})
	processor["detectors"] = []string{"env", "system", "ec2", "docker", "elastic_beanstalk", "lambda"}

	return processor
}

func marshalResourceProcessor(_ context.Context, telemetry *config.Telemetry) map[string]interface{} {
	processor := make(map[string]interface{})
	processor["attributes"] = []struct{ Key, Value, Action string }{
		{"service.instance.id", telemetry.Resource.GetID(), "upsert"},
		{"service.namespace", telemetry.Resource.GetNamespace(), "upsert"},
		{"service.version", telemetry.Resource.GetVersion(), "upsert"},
		{"service.name", telemetry.Resource.GetName(), "upsert"},
	}

	return processor
}
