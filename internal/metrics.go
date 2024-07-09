package internal

import "time"

const (
	// DefaultMetricReadInterval is the default interval for reading metricsold.
	// It is set to 5 seconds.
	DefaultMetricReadInterval = 5 * time.Second

	// DefaultMetricExportInterval is the default interval for exporting metricsold.
	// It is set to 5 seconds.
	DefaultMetricExportInterval = 5 * time.Second
)
