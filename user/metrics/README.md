# :chart_with_downwards_trend: Metrics

## Getting started

We use [Prometheus](https://prometheus.io/) for sending metrics.
To start to instrument your application, you can use [this Prometheus Go client](https://github.com/prometheus/client_golang).
You can create a `metrics.go` file under `metrics`.
An example code would be:

```go
package metrics

import (
 "github.com/prometheus/client_golang/prometheus"
 "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
 ExampleCount = promauto.NewCounter(
  prometheus.CounterOpts{
   Name: "service_metrics",
   Help: "The total number of something",
  },
 )
)
```

You can see the added metric on the `metrics/` endpoint that is exposed from `server.go`.

You can find our common metrics in `common/go/metrics`.

For more information and examples, take a look at the [official docs](https://prometheus.io/docs/guides/go-application/)
