package errors

import "errors"

var ErrTelemetryNotEnabled = errors.New("telemetry is not enabled")

var ErrTelemetryMetricsNotEnabled = errors.New("telemetry metrics is not enabled")
var ErrTelemetryLogsNotEnabled = errors.New("telemetry logs is not enabled")
var ErrTelemetryTracesNotEnabled = errors.New("telemetry traces is not enabled")
