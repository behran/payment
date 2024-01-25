package metric

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/fx"
)

var (
	//ResponseTime время ответа сервиса по http ....
	ResponseTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration",
		Help:    "Response duration in seconds.",
		Buckets: prometheus.DefBuckets,
	}, []string{"path", "method", "code"})
	//CodeHTTPCounter кол-во запросов ...
	CodeHTTPCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of requests by status code.",
	}, []string{"path", "method", "code"})
)

// RegistrationMetrics ...
func RegistrationMetrics(lc fx.Lifecycle) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			prometheus.MustRegister(ResponseTime, CodeHTTPCounter)
			return nil
		},
	})
	return nil
}
