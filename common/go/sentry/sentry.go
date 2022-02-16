package sentry

import (
	"os"

	"github.com/getsentry/sentry-go"
)

func fillEnvironment(opts sentry.ClientOptions) sentry.ClientOptions {
	if opts.Environment == "" {
		env := os.Getenv("ENV")

		if env != "" {
			opts.Environment = env
		}
	}

	return opts
}

// NewClient creates a new Sentry client with your options.
// If you fail to specify some of them, it turns them into our defaults.
// By default, as well, from Sentry package:
// If not specified in the SDK initialization, the DSN, Release and Environment are read from
// the environment variables SENTRY_DSN, SENTRY_RELEASE and SENTRY_ENVIRONMENT, respectively.
// However, we'll also attempt to obtain SENTRY_ENVIRONMENT from ENV since that is our normal
// environment variable for environment.
func NewClient(opts sentry.ClientOptions) (*sentry.Client, error) {
	opts = fillEnvironment(opts)
	return sentry.NewClient(opts)
}

// Init initialises the Sentry SDK with your options.
// If you fail to specify some of them, it turns them into our defaults.
// By default, as well, from Sentry package: If not specified in the SDK initialization,
// the DSN, Release and Environment are read from the environment variables SENTRY_DSN,
// SENTRY_RELEASE and SENTRY_ENVIRONMENT, respectively.
// However, we'll also attempt to obtain SENTRY_ENVIRONMENT from ENV since
// that is our normal environment variable for environment.
func Init(opts sentry.ClientOptions) error {
	opts = fillEnvironment(opts)
	return sentry.Init(opts)
}
