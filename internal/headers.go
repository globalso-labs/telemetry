package internal

import (
	"maps"

	"go.globalso.dev/x/telemetry/common"
	"go.globalso.dev/x/telemetry/internal/constants"
)

var defaultHeaders = map[string]string{
	constants.HeaderTracer: "OpenTelemetry",
}

func GetHeaders() map[string]string {
	var headers = make(map[string]string)
	maps.Copy(headers, defaultHeaders)

	headers[constants.HeaderScopeOrgID] = common.OrganizationID()

	return headers
}
