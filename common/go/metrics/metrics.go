package metrics

import (
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/crafted-systems/smartcollect-pro/common/go/sort"
)

var (
	// DefBuckets are the default buckets we use for our services.
	DefBuckets = sort.Float64s(append([]float64{.001, .003, .75, 1.25, 1.5, 1.75}, prometheus.DefBuckets...))

	upstreamRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "upstream_request_duration_seconds",
		Help:    "The time it took to complete the request upstream",
		Buckets: DefBuckets,
	},
		[]string{"service", "status_code", "endpoint", "upstream_service"},
	)

	requestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "request_duration_total_seconds",
			Help: "The duration of a request",
		},
		[]string{"service", "method", "endpoint", "status_code"},
	)
)

// TrackUpstreamRequestDuration refers to the metric upstream_request_duration_seconds.
// This tracks a request made upstream (i.e. to a different service).
func TrackUpstreamRequestDuration(duration time.Duration, service, upstreamService, statusCode, endpoint string) {
	upstreamRequestDuration.With(
		prometheus.Labels{
			"service":          service,
			"status_code":      statusCode,
			"endpoint":         endpoint,
			"upstream_service": upstreamService,
		},
	).Observe(duration.Seconds())
}

// TrackRequestDuration tracks the duration of a request made to the service.
// This is used only for http requests, as grpc requests use the default server integration.
func TrackRequestDuration(duration time.Duration, service, method, statusCode, endpoint string) {
	requestDuration.With(
		prometheus.Labels{
			"service":     service,
			"status_code": statusCode,
			"endpoint":    endpoint,
			"method":      method,
		},
	).Observe(duration.Seconds())
}

// StartServer starts a generic metric server.
func StartServer(path, address string) {
	http.Handle(path, promhttp.Handler())

	go func() {
		err := http.ListenAndServe(address, nil)
		if err != nil {
			zap.L().Error("metrics server stopped", zap.Error(err))
		}
	}()
}
