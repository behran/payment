package middlewares

import (
	"github.com/valyala/fasthttp"
)

// Middleware ...
type Middleware func(fasthttp.RequestHandler) fasthttp.RequestHandler

const (
	PayloadKey   = "payload"
	AccountKey   = "account"
	AccountIDKey = "id"
)

var (
	CreateAccountPayload = []Middleware{
		HeadersMiddleware,
		AccountMiddleware,
		MetricTimeResponseMiddleware,
	}
	UpdateAccountPayload = []Middleware{
		PayloadMiddleware,
		CheckValidHeadersMiddleware,
		AccountIDMiddleware,
		HeadersMiddleware,
		MetricTimeResponseMiddleware,
	}
)

// ApplyMiddleware iterator of middleware ...
func ApplyMiddleware(handle fasthttp.RequestHandler, middleware ...Middleware) fasthttp.RequestHandler {
	if len(middleware) < 1 {
		return handle
	}
	wrapped := handle

	for i := len(middleware) - 1; i >= 0; i-- {
		wrapped = middleware[i](wrapped)
	}
	return wrapped
}
