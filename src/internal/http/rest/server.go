package rest

import (
	"context"
	"net"

	"payment/internal/config"

	"github.com/buaazp/fasthttprouter"
	"github.com/lab259/cors"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

//StartServer ...
func StartServer(lc fx.Lifecycle, logger *zap.Logger, config config.Config, router *fasthttprouter.Router) error {
	addr := net.JoinHostPort("", config.App.Port)
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Listening http server ...\n", zap.String("addr", addr))
			return fasthttp.ListenAndServe(
				addr,
				cors.AllowAll().Handler(router.Handler), // append cors ...
			)
		},
	})
	return nil
}
