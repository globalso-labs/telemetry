/*
 * telemetry
 * telemetry.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Thu, 12 Sep 2024 21:16:25 -0500 by nick.
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
	"go.globalso.dev/x/telemetry/common"
	"gopkg.in/yaml.v2"
)

type Telemetry struct {
	Enabled  bool    `yaml:"enabled"`
	Endpoint string  `yaml:"endpoint"`
	Protocol string  `yaml:"protocol"`
	Headers  Headers `yaml:"headers"`
	Scrape   Scrape  `yaml:"scrape"`
	Push     Push    `yaml:"push"`
	Agent    Agent   `yaml:"agent"`
	Logger   Logger  `yaml:"logger"`
	Meter    Meter   `yaml:"meter"`
	Tracer   Tracer  `yaml:"tracer"`

	Resource *common.Resource `yaml:"resource"`
}

// Enable activates the telemetry and its associated components (Agent, Logger, Meter, Tracer).
func (t *Telemetry) Enable() {
	t.Enabled = true
	t.Logger.Enabled = true
	t.Meter.Enabled = true
	t.Tracer.Enabled = true
}

// Disable deactivates the telemetry and its associated components (Tracer, Meter, Logger, Agent).
func (t *Telemetry) Disable() {
	t.Tracer.Enabled = false
	t.Meter.Enabled = false
	t.Logger.Enabled = false
	t.Enabled = false
}

func (t *Telemetry) Dump() ([]byte, error) {
	return yaml.Marshal(t)
}
