package application

import (
	"context"

	"payment/internal/config"
	"payment/internal/database"
	"payment/internal/demon"
	"payment/internal/factory"
	"payment/internal/http/rest"
	"payment/pkg/log"

	"github.com/buaazp/fasthttprouter"
	"go.uber.org/fx"
)

//Run DI containers ...
func Run() error {
	containers := fx.New(
		fx.Provide(
			config.New,
			fasthttprouter.New,
			logger.New,
			database.NewConnectManager,
			factory.NewFactory,
			demon.New,
		),
		fx.Invoke(
			database.InitConnections, // init connection database ...
			factory.InitServices,     // create domain services ...
			demon.Start,              // start event rollback tx ...
			rest.RegisterRoutes,      // registration handle routes ...
			rest.StartServer,         // start http server ...
		),
	)
	return containers.Start(context.Background())
}
