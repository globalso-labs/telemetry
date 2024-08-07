/*
 * telemetry
 * tracer.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 6 Aug 2024 22:11:00 -0500 by nick.
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

package tracer

import (
	"sync"

	"go.globalso.dev/x/telemetry/internal"
	"go.opentelemetry.io/otel/trace"
)

// _tracer is a singleton instance of a trace.Tracer.
// It is initialized only once and used throughout the application.
var _tracer trace.Tracer

// _sync is a mutex used to synchronize the initialization of the _tracer instance.
var _sync sync.Mutex

// Tracer returns a singleton instance of a trace.Tracer.
// It ensures that the tracer is initialized only once using a mutex for synchronization.
//
// Returns:
// - trace.Tracer: The singleton tracer instance.
func Tracer() trace.Tracer { //nolint:ireturn
	_sync.Lock()
	defer _sync.Unlock()

	if _tracer == nil {
		_tracer = _handler.Provider().Tracer(
			internal.Module,
			trace.WithInstrumentationVersion(internal.Version),
		)
	}

	return _tracer
}
