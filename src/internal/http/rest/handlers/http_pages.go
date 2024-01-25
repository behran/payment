package handlers

import (
	"payment/internal/http/rest/response"
	"payment/internal/http/rest/response/errors"
	logger "payment/pkg/log"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

// PageNotFound ...
func PageNotFound(ctx *fasthttp.RequestCtx) { response.Error(ctx, errors.ErrNotFound) }

// PagePanic ...
func PagePanic(ctx *fasthttp.RequestCtx, i any) {
	logger.Logger.Error("error handle panic", zap.Any("panic", i))

	response.Error(ctx, errors.ErrServer)
}
