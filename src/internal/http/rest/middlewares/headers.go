package middlewares

import (
	"payment/internal/dto"
	"payment/internal/http/rest/response"
	"payment/internal/http/rest/response/errors"
	"payment/internal/tools"

	"github.com/valyala/fasthttp"
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

//HeadersMiddleware ...
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

//CheckValidHeadersMiddleware ...
func CheckValidHeadersMiddleware(handle fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		payload := ctx.UserValue(PayloadKey).(dto.Payload)
		sourceType := string(ctx.Request.Header.Peek("Source-Type"))
		payload.SetSourceType(sourceType)

		if !tools.IsExistSlice(payload.SourceType(), []string{
			sourceGame, sourcePayment, sourceServer,
		}) {
			response.Error(errors.ErrSourceType, ctx)
			return
		}

		ctx.SetUserValue(PayloadKey, payload)

		// next
		handle(ctx)
	}
}
