package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecification_Parse(t *testing.T) {
	os.Setenv("PORT", "9001")
	os.Setenv("METRICS_PORT", "9002")
	os.Setenv("ENV", "iire9457fjfjjkfjkf")
	os.Setenv("LOG_LEVEL", "kdkoe8499599589jfjf")

	spec, err := Parse()

	if assert.Nil(t, err) {
		assert.Equal(t, spec.Port, "9001", fmt.Sprintf("Expected Port %v, Got Port %v", 9001, spec.Port))
		assert.Equal(t, spec.MetricsPort, "9002", fmt.Sprintf("Expected Port %v, Got Port %v", 9002, spec.MetricsPort))
		assert.Equal(t, spec.Env, "iire9457fjfjjkfjkf", fmt.Sprintf("Expected Port %v, Got Port %v", "iire9457fjfjjkfjkf", spec.Env))
		assert.Equal(t, spec.LogLevel, "kdkoe8499599589jfjf", fmt.Sprintf("Expected Log Level %v, Got Log Level %v", "kdkoe8499599589jfjf", spec.LogLevel))
	}
}
