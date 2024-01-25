package response

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"payment/internal/contracts"
	logger "payment/pkg/log"
	"payment/pkg/metric"

	"github.com/valyala/fasthttp"
)

// Response ...
type Response struct {
	Meta     Meta    `json:"meta"`
	JSONAPI  JSONAPI `json:"jsonapi"`
	Response any     `json:"data,omitempty"`
	Errors   []Err   `json:"errors,omitempty"`
}

// Response objects ...
type (
	//Meta ...
	Meta struct {
		Status bool `json:"status"`
		Code   int  `json:"status_code"`
	}
	//JSONAPI ...
	JSONAPI struct {
		Version string `json:"version"`
	}
	//Err ...
	Err struct {
		Tittle string `json:"title"`
		Detail string `json:"detail"`
	}
)

// Version ...
var Version = os.Getenv("API_VERSION")

// OK ...
func OK(ctx *fasthttp.RequestCtx, body any) {
	var payload = Response{
		Meta: Meta{
			Code:   http.StatusOK,
			Status: true,
		},
		JSONAPI: JSONAPI{
			Version: Version,
		},
		Response: body,
	}
	send(ctx, payload, http.StatusOK)
}

// Error ...
func Error(ctx *fasthttp.RequestCtx, err error) {
	e := err.(contracts.IPaymentError)
	payload := Response{
		Meta: Meta{
			Code:   e.StatusCode(),
			Status: false,
		},
		JSONAPI: JSONAPI{
			Version: Version,
		},
		Errors: []Err{
			{
				Tittle: e.Error(),
				Detail: e.Details(),
			},
		},
	}
	send(ctx, payload, e.StatusCode())
}

func send(ctx *fasthttp.RequestCtx, r Response, code int) {
	metric.CodeHTTPCounter.With(prometheus.Labels{
		"code":   strconv.Itoa(code),
		"method": string(ctx.Method()),
		"path":   string(ctx.Path()),
	}).Inc()
	ctx.SetStatusCode(code)

	if err := json.NewEncoder(ctx).Encode(r); err != nil {
		logger.Logger.Error("fail send response", zap.Error(err))
	}
}
