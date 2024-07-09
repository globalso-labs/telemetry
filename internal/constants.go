package internal

const (
	// TelemetryEndpoint is the endpoint for telemetry data.
	TelemetryEndpoint = "telemetry.idbi.pe"

	// TelemetryMetricsPath is the URL path for sending telemetry metricsold.
	TelemetryMetricsPath = "otlp/v1/metricsold"

	// TelemetryTracesPath is the URL path for sending telemetry traces.
	TelemetryTracesPath = "otlp/v1/traces"

	// TelemetryLogsPath is the URL path for sending telemetry logs.
	TelemetryLogsPath = "otlp/v1/logs"
)
