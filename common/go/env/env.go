package env

import (
	"fmt"
	"os"
)

type Environment string

const (
	EnvironmentProd          = "prod"
	EnvironmentStaging       = "staging"
	EnvironmentBranchPreview = "branch"
	EnvironmentDev           = "dev"
)

var currentEnv Environment

func Parse() error {
	switch env := os.Getenv("ENV"); env {
	case EnvironmentProd:
		currentEnv = EnvironmentProd
	case EnvironmentStaging:
		currentEnv = EnvironmentStaging
	case EnvironmentBranchPreview:
		currentEnv = EnvironmentBranchPreview
	case EnvironmentDev:
		currentEnv = EnvironmentDev
	default:
		return fmt.Errorf("unable to parse ENV: %s", env)
	}
	return nil
}

func IsProd() bool {
	return currentEnv == EnvironmentProd
}

func IsStaging() bool {
	return currentEnv == EnvironmentStaging
}

func IsBranchPreview() bool {
	return currentEnv == EnvironmentBranchPreview
}

func IsDev() bool {
	return currentEnv == EnvironmentDev
}
