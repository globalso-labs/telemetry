/*
 * telemetry
 * parse.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Sat, 14 Sep 2024 21:35:27 -0500 by nick.
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
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// GetResource returns the service resource.
func GetResource(ctx context.Context, res *Resource) *resource.Resource {
	attrs := make([]attribute.KeyValue, 0)
	attrs = append(attrs, semconv.ServiceInstanceIDKey.String(res.GetID()))
	attrs = append(attrs, semconv.ServiceNameKey.String(res.GetName()))
	attrs = append(attrs, semconv.ServiceNamespaceKey.String(res.GetNamespace()))
	attrs = append(attrs, semconv.ServiceVersionKey.String(res.GetVersion()))

	output, err := resource.New(ctx,
		resource.WithSchemaURL(semconv.SchemaURL),
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithOS(),
		resource.WithHost(),
		resource.WithHostID(),
		resource.WithAttributes(attrs...),
	)

	if err != nil {
		return resource.NewWithAttributes(semconv.SchemaURL, attrs...)
	}

	return output
}
