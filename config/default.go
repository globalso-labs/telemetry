/*
 * telemetry
 * parse.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Thu, 12 Sep 2024 22:02:08 -0500 by nick.
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

package config

import (
	"go.globalso.dev/x/telemetry/constants"
	"go.globalso.dev/x/telemetry/internal"
)

// Default initializes and returns a default Telemetry configuration.
// The configuration includes default values for various telemetry settings
// such as Endpoint, Protocol, Headers, Scrape, Push, Agent, Logger, Meter, and Tracer.
func Default() *Telemetry {
	t := &Telemetry{
		Enabled:  true,
		Endpoint: constants.Endpoint,
		Protocol: constants.Protocol,
		Headers:  Headers{},
		Scrape: Scrape{
			Interval: constants.ScrapeInterval,
		},
		Push: Push{
			Interval: constants.PushInterval,
		},
		Agent: Agent{
			Extensions: map[string]interface{}{},
			Exporters:  map[string]interface{}{},
			Receivers:  map[string]interface{}{},
			Processors: map[string]interface{}{},
			Service:    map[string]interface{}{},
		},
		Logger: LoggerDefault(),
		Meter:  MeterDefault(),
		Tracer: TracerDefault(),

		Resource: internal.NewResource(),
	}

	t.Enable()
	return t
}
