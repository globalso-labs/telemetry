/*
 * telemetry
 * telemetry_test.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Wed, 31 Jul 2024 20:30:53 -0500 by nick.
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

package telemetry_test

import (
	"context"
	"io"
	"math/rand"
	"testing"
	"time"

	"go.globalso.dev/x/telemetry"
	"go.globalso.dev/x/telemetry/common"
	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/logger"
	"go.globalso.dev/x/telemetry/logger/constants"
)

func Test_Telemetry(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	lOpts := []logger.Option{
		logger.WithLevel(constants.TraceLevel),
		logger.WithWriter(io.Discard),
	}

	cOpts := []common.Option{
		common.WithVersion("1.0.0"),
	}
	cfg := config.New(
		config.WithLoggerOpts(lOpts...),
		config.WithCommonOpts(cOpts...),
	)
	telemetry.Execute(ctx, cfg)

	// We aren't testing the Fatal constants here, since the test will exit after the first call to Fatal.
	for {
		select {
		case <-ctx.Done():
			return
		default:
			switch rand.Intn(5) {
			case 0:
				logger.Trace().Msg("This is a trace message")
			case 1:
				logger.Debug().Msg("This is a debug message")
			case 2:
				logger.Info().Msg("This is an info message")
			case 3:
				logger.Warn().Msg("This is a warn message")
			case 4:
				logger.Error().Msg("This is an error message")
			}
		}
	}
}
