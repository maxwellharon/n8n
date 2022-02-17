package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Specification struct to load environment variables.
type Specification struct {
	MetricsPort    string `default:"9099" split_words:"true"`
	Port           string `default:"8888"`
	GRPCReflection bool   `default:"false"`
	Env            string `default:"branch" split_words:"true"`
	LogLevel       string `default:"debug" split_words:"true"`
}

// Parse reads config from environment.
func Parse() (Specification, error) {
	config := Specification{}

	err := envconfig.Process("", &config)
	if err != nil {
		return config, fmt.Errorf("reading config: %w", err)
	}

	return config, nil
}
