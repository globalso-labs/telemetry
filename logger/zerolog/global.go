/*
 * telemetry
 * global.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Wed, 31 Jul 2024 15:49:42 -0500 by nick.
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

package zerolog

import (
	"github.com/rs/zerolog"
	"go.globalso.dev/x/telemetry/internal/constants"
	"go.globalso.dev/x/telemetry/logger/level"
)

var DefaultContextLogger = zerolog.New(NewConsoleWriter()).
	Level(FromLevel(constants.DefaultLoggerLevel)).With().
	Timestamp().
	Caller().
	Logger()

func FromLevel(l level.Level) zerolog.Level {
	switch l {
	case level.TraceLevel:
		return zerolog.TraceLevel
	case level.DebugLevel:
		return zerolog.DebugLevel
	case level.InfoLevel:
		return zerolog.InfoLevel
	case level.WarnLevel:
		return zerolog.WarnLevel
	case level.ErrorLevel:
		return zerolog.ErrorLevel
	case level.FatalLevel:
		return zerolog.FatalLevel
	}
	return zerolog.NoLevel
}
