package middlewares

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"
	"payment/pkg/metric"
)

// MetricTimeResponseMiddleware обработчик времени выполнения запроса ...
func MetricTimeResponseMiddleware(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()
		handler(ctx)
		// diff time to prometheus metrics ...
		metric.ResponseTime.With(prometheus.Labels{
			"method": string(ctx.Method()),
			"path":   string(ctx.Path()),
			"code":   strconv.Itoa(ctx.Response.StatusCode()),
		}).Observe(time.Since(start).Seconds())
	}
}
