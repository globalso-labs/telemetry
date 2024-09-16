/*
 * telemetry
 * options.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Sat, 14 Sep 2024 21:40:20 -0500 by nick.
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

package telemetry

import (
	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/internal"
	"go.globalso.dev/x/telemetry/logger"
)

// Option represents a function that applies a configuration to the Telemetry instance.
type Option func(*Telemetry)

// WithLogger sets the logger for the Telemetry instance.
//
// Parameters:
// - logger *logger.Logger: The logger instance to be used.
//
// Returns:
// - Option: A function that sets the logger for the Telemetry instance.
func WithLogger(logger *logger.Logger) Option {
	return func(t *Telemetry) {
		t.logger = logger
	}
}

// WithConfig sets the configuration for the Telemetry instance.
//
// Parameters:
// - config *config.Telemetry: The configuration to be used.
//
// Returns:
// - Option: A function that sets the configuration for the Telemetry instance.
func WithConfig(config *config.Telemetry) Option {
	return func(t *Telemetry) {
		t.config = config
	}
}

// WithResource sets the resource for the Telemetry instance.
//
// Parameters:
// - resource *internal.Resource: The resource to be used.
//
// Returns:
// - Option: A function that sets the resource for the Telemetry instance.
func WithResource(resource *internal.Resource) Option {
	return func(t *Telemetry) {
		t.resource = resource
	}
}
