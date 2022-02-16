module github.com/crafted-systems/smartcollect-pro/common/go/logger

go 1.17

replace github.com/crafted-systems/smartcollect-pro/common/go/sentry v0.0.0 => ../sentry

replace github.com/crafted-systems/smartcollect-pro/common/go v0.0.0 => ../

require (
	github.com/TheZeroSlave/zapsentry v1.9.0
	github.com/crafted-systems/smartcollect-pro/common/go v0.0.0
	github.com/crafted-systems/smartcollect-pro/common/go/sentry v0.0.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.21.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/getsentry/sentry-go v0.11.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
