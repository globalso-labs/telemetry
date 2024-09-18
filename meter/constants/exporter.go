/*
 * telemetry
 * exporter.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Sun, 4 Aug 2024 00:07:28 -0500 by nick.
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

package constants

import "time"

const (
	// DefaultMetricReadInterval is the default interval for reading metrics.
	// It is set to 5 seconds.
	DefaultMetricReadInterval = 5 * time.Second

	// DefaultMetricExportInterval is the default interval for exporting metrics.
	// It is set to 5 seconds.
	DefaultMetricExportInterval = 5 * time.Second
)