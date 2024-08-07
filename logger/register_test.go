/*
 * telemetry
 * register_test.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Wed, 31 Jul 2024 15:22:38 -0500 by nick.
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

package logger_test

import (
	"context"
	"io"
	"math/rand"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"go.globalso.dev/x/telemetry/logger"
)

func Test_SendLogs(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := []logger.Option{
		logger.WithLevel(zerolog.TraceLevel),
		logger.WithWriter(io.Discard),
	}
	cfg := logger.NewConfig(opts...)
	err := logger.Register(ctx, &cfg)
	assert.Nil(t, err)
	defer logger.Shutdown(ctx)

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
