module github.com/crafted-systems/smartcollect-pro/user

go 1.17

require (
	github.com/crafted-systems/smartcollect-pro/common/go v0.0.0
	github.com/crafted-systems/smartcollect-pro/common/go/logger v0.0.0
	github.com/golang/mock v1.6.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.21.0
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.26.0
)

require (
	github.com/TheZeroSlave/zapsentry v1.9.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/crafted-systems/smartcollect-pro/common/go/sentry v0.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/getsentry/sentry-go v0.11.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.12.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200825200019-8632dd797987 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/crafted-systems/smartcollect-pro/common/go v0.0.0 => ../common/go
	github.com/crafted-systems/smartcollect-pro/common/go/logger v0.0.0 => ../common/go/logger
	github.com/crafted-systems/smartcollect-pro/common/go/sentry v0.0.0 => ../common/go/sentry
	github.com/crafted-systems/smartcollect-pro/common/proto v0.0.0 => ../common/proto
)
