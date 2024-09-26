/*
 * telemetry
 * sentry.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Sun, 15 Sep 2024 18:37:22 -0500 by nick.
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

package hooks

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

const sentryFlushTimeout = 2 * time.Second

type SentryHook struct{}

func (h SentryHook) Run(event *zerolog.Event, level zerolog.Level, message string) {
	if level >= zerolog.ErrorLevel && level <= zerolog.PanicLevel {
		captured := h.convertEvent(event, level, message)
		sentry.CaptureEvent(&captured)
	}

	if level == zerolog.FatalLevel || level == zerolog.PanicLevel {
		sentry.Flush(sentryFlushTimeout)
	}
}

func (h SentryHook) convertEvent(e *zerolog.Event, level zerolog.Level, message string) sentry.Event {
	var record sentry.Event

	record.Level = sentry.Level(level.String())
	record.Message = message
	record.Timestamp = zerolog.TimestampFunc()
	record.Extra = convertSentryFields(e)
	return record
}

func convertSentryFields(e *zerolog.Event) map[string]interface{} {
	kvs := make(map[string]interface{})

	// Extract fields from the event and convert them
	e.Fields(func(key string, value interface{}) {
		kvs[key] = value
	})

	return kvs
}
