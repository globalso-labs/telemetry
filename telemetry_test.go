package telemetry_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.globalso.dev/x/telemetry/config"
)

var ctx = context.Background()

func Test_Telemetry(t *testing.T) {
	cfg := config.New()
	assert.NotNil(t, cfg)
}
