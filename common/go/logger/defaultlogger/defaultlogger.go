package defaultlogger

import (
	"fmt"
	"github.com/TheZeroSlave/zapsentry"
	"github.com/crafted-systems/smartcollect-pro/common/go/errors"
	"github.com/crafted-systems/smartcollect-pro/common/go/logger"
	custom_sentry "github.com/crafted-systems/smartcollect-pro/common/go/sentry"
	custom_zap_sentry "github.com/crafted-systems/smartcollect-pro/common/go/sentry/zap"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
)

// Options are our options for the default logger.
// These are different from the normal logger in the sense that they are a subset.
// Most likely in the future they will remain a subset, but they can also be shortcuts
// to setting several options (eg. a flag called "sentryEnabled" that
// sets and creates a core, or other similar stuff).
type Options struct {
	ignoredErrors []error
}

// Option are functional options:
// You can add options without changing the interface of your service.
// Eg: https://github.com/tmrts/go-patterns/blob/master/idiom/functional-options.md.
type Option func(*Options)

// NewLogger creates our default logger: it uses a NewDefaultConfig(level) as a base and adds a sentry core.
func NewLogger(level string, options ...Option) (*zap.Logger, error) {
	// apply the options.
	var opts Options
	for _, o := range options {
		o(&opts)
	}

	// if we have no ignored errors option, we default to ignoring the common error "ErrExpected".
	if len(opts.ignoredErrors) == 0 {
		opts.ignoredErrors = []error{errors.ErrExpected}
	}

	logCfg, err := logger.NewDefaultConfig(level)
	if err != nil {
		fmt.Printf("warning: creating config: %v\n", err)
	}

	// we default to creating a sentry core, since all our services, to date, use Sentry.
	client, err := custom_sentry.NewClient(sentry.ClientOptions{AttachStacktrace: true})
	if err != nil {
		return nil, fmt.Errorf("creating sentry client: %w", err)
	}

	sentryCore, err := zapsentry.NewCore(
		custom_zap_sentry.DefaultConfiguration(),
		zapsentry.NewSentryClientFromClient(client),
	)
	if err != nil {
		return nil, fmt.Errorf("creating sentry core: %w", err)
	}

	log, err := logger.NewLogger(
		logCfg,
		logger.WithExtraCores(sentryCore),
		logger.WithIgnoredErrors(opts.ignoredErrors...),
	)
	if err != nil {
		return nil, fmt.Errorf("creating logger: %w", err)
	}

	return log, nil
}

// WithIgnoredErrors ignores writing certain errors.
func WithIgnoredErrors(errs []error) Option {
	return func(o *Options) {
		o.ignoredErrors = errs
	}
}
