package rest

import (
	"context"
	"net"

	"payment/internal/config"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Server ...
type Server struct {
	http *fasthttp.Server
	addr string
}

// NewServer ...
func NewServer(config config.Config, router *router.Router) *Server {
	return &Server{
		addr: net.JoinHostPort(config.App.Hostname, config.App.Port),
		http: &fasthttp.Server{
			Handler: router.Handler,
		},
	}
}

// StartServer ...
func StartServer(lc fx.Lifecycle, logger *zap.Logger, server *Server) error {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Listening http server ...\n", zap.String("addr", server.addr))
			go func(s *Server) {
				if err := server.http.ListenAndServe(s.addr); err != nil {
					logger.Error("fail start http server",
						zap.String("addr", s.addr),
						zap.Error(err),
					)
				}
			}(server)
			return nil
		},
		OnStop: func(context.Context) error {
			return server.http.Shutdown()
		},
	})
	return nil
}
