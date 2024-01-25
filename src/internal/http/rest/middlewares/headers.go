package middlewares

import (
	"slices"

	"github.com/valyala/fasthttp"
	"payment/internal/dto"
	"payment/internal/http/rest/response"
	"payment/internal/http/rest/response/errors"
)

const (
	contentType          = "application/json"
	corsAllowHeaders     = "authorization"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

const (
	sourceGame    = "game"
	sourceServer  = "server"
	sourcePayment = "payment"
)

// HeadersMiddleware ...
func HeadersMiddleware(handle fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set(fasthttp.HeaderContentType, contentType)
		ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowCredentials, corsAllowCredentials)
		ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowHeaders, corsAllowHeaders)
		ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowMethods, corsAllowMethods)
		ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowOrigin, corsAllowOrigin)
		// next
		handle(ctx)
	}
}

// CheckValidHeadersMiddleware ...
func CheckValidHeadersMiddleware(handle fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		sourceType := string(ctx.Request.Header.Peek("Source-Type"))

		payload := ctx.UserValue(PayloadKey).(dto.Payload)
		payload.SetSourceType(sourceType)

		if !slices.Contains([]string{
			sourceGame, sourcePayment, sourceServer,
		}, payload.SourceType()) {
			response.Error(ctx, errors.ErrSourceType)
			return
		}
		ctx.SetUserValue(PayloadKey, payload)
		// next
		handle(ctx)
	}
}
