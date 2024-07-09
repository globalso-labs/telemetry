/*
 * telemetry
 * common.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Mon, 8 Jul 2024 21:05:00 -0500 by nick.
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

package common

// Common holds configuration settings for telemetry.
// It includes the organization ID, OTLP endpoint, and OTLP path.
type Common struct {
	ID        string // ID is the unique identifier for the telemetry instance. Used to identify the running instance.
	Name      string // Name is the name of the service or application. Used to identify the service or application.
	Namespace string // Namespace is the name of the group of services. Used to group services or applications.
	Version   string // Version is the version of the service. Used to identify the version of service.

	OrganizationID string // OrganizationID is the unique identifier for the organization.
	OTLPEndpoint   string // OTLPEndpoint is the endpoint for OpenTelemetry Protocol (OTLP) communication.

}
