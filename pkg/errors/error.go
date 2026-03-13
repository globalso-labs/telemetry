package errors

import "errors"

var ErrTelemetryNotEnabled = errors.New("telemetry is not enabled")
var ErrTelemetryConfigNil = errors.New("telemetry config is nil")

var ErrTelemetryMeterNotEnabled = errors.New("telemetry metrics is not enabled")
var ErrTelemetryLoggerNotEnabled = errors.New("telemetry logs is not enabled")
var ErrTelemetryTracerNotEnabled = errors.New("telemetry traces is not enabled")
