/*
 * telemetry
 * resource.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 9 Jul 2024 00:25:06 -0500 by nick.
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

package internal

import (
	"go.globalso.dev/x/telemetry/common"
	"go.opentelemetry.io/otel/sdk/resource"

	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// GetResource returns the service resource.
func GetResource() *resource.Resource {
	attrs := resource.Default().Attributes()
	attrs = append(attrs, semconv.ServiceInstanceIDKey.String(common.ID()))
	attrs = append(attrs, semconv.ServiceNameKey.String(common.Name()))
	attrs = append(attrs, semconv.ServiceNamespaceKey.String(common.Namespace()))
	attrs = append(attrs, semconv.ServiceVersionKey.String(common.Version()))

	return resource.NewWithAttributes(semconv.SchemaURL, attrs...)
}
