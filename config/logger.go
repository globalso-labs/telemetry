/*
 * telemetry
 * logger.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Thu, 12 Sep 2024 21:18:48 -0500 by nick.
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
	"gopkg.in/yaml.v3"
)

type Logger struct {
	Enabled bool     `yaml:"enabled"`
	Path    string   `yaml:"path"`
	Level   string   `yaml:"level"`
	Drivers []string `yaml:"drivers"`
	Hooks   []string `yaml:"hooks"`
	Push    Push     `yaml:"push"`
}

func (l *Logger) Enable() {
	l.Enabled = true
}

func (l *Logger) Disable() {
	l.Enabled = false
}

func (l *Logger) Dump() ([]byte, error) {
	return yaml.Marshal(l)
}

func LoggerDefault() Logger {
	return Logger{
		Enabled: true,
		Path:    constants.LoggerPath,
		Level:   "trace",
		Drivers: []string{"console"},
		Hooks:   []string{"otlp"},
		Push: Push{
			Interval: constants.PushInterval,
		},
	}
}
