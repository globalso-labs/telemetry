/*
 * telemetry
 * receivers.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 24 Sep 2024 03:30:43 -0500 by nick.
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

// marshalReceivers marshals the telemetry receivers configuration into a map.
//
// Parameters:
// - ctx context.Context: The context to use for marshaling.
// - telemetry *config.Telemetry: The telemetry configuration to be marshaled.
//
// Returns:
// - map[string]interface{}: The marshaled receivers configuration.
// - error: An error if the marshaling fails.
func marshalReceivers(ctx context.Context, telemetry *config.Telemetry) map[string]interface{} {
	receivers := make(map[string]interface{})

	receivers["nop"] = marshalNOPReceiver(ctx, telemetry)

	return receivers
}

// marshalNOPReceiver marshals the NOP receiver configuration into a map.
//
// Parameters:
// - ctx context.Context: The context to use for marshaling (unused).
// - telemetry *config.Telemetry: The telemetry configuration to be marshaled (unused).
//
// Returns:
// - map[string]interface{}: The marshaled NOP receiver configuration.
func marshalNOPReceiver(_ context.Context, _ *config.Telemetry) map[string]interface{} {
	return nil
}
