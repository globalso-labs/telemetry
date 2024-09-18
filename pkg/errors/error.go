package errors

import "errors"

var ErrTelemetryNotEnabled = errors.New("telemetry is not enabled")

var ErrTelemetryMeterNotEnabled = errors.New("telemetry metrics is not enabled")
var ErrTelemetryLoggerNotEnabled = errors.New("telemetry logs is not enabled")
var ErrTelemetryTracerNotEnabled = errors.New("telemetry traces is not enabled")
