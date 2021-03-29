package middlewares

import (
	"time"

	logger "payment/pkg/log"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LogMiddleware(handle fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		startTime := time.Now()
		handle(ctx)
		var addrField zapcore.Field
		xRealIp := ctx.Request.Header.Peek("X-Real-IP")
		if len(xRealIp) > 0 {
			addrField = zap.ByteString("addr", ctx.Request.Header.Peek("X-Real-IP"))
		} else {
			addrField = zap.String("addr", ctx.RemoteAddr().String())
		}

		logger.Logger.Info("access",
			zap.Int("code", ctx.Response.StatusCode()),
			zap.Duration("time", time.Since(startTime)),
			zap.ByteString("method", ctx.Method()),
			zap.ByteString("path", ctx.Path()),
			zap.ByteString("agent", ctx.UserAgent()),
			zap.ByteString("req", ctx.RequestURI()),
			addrField)
	}
}
