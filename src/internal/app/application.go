package application

import (
	"payment/internal/config"
	"payment/internal/database"
	"payment/internal/demon"
	"payment/internal/factory"
	"payment/internal/http/rest"
	"payment/pkg/log"
	"payment/pkg/metric"

	"github.com/fasthttp/router"
	"go.uber.org/fx"
)

// Containers DI containers ...
func Containers() *fx.App {
	return fx.New(
		fx.Provide(
			config.New,
			router.New,
			logger.New,
			rest.NewServer,
			database.NewConnectManager,
			factory.NewFactory,
			demon.New,
		),
		fx.Invoke(
			database.InitConnections,   // init connection database ...
			factory.InitServices,       // create domain services ...
			demon.Start,                // start event rollback tx ...
			rest.RegisterRoutes,        // registration handle routes ...
			rest.StartServer,           // start http server ...
			metric.RegistrationMetrics, // register prometheus metrics ...
		),
	)
}

// Cli DI containers for cli command ...
func Cli() *fx.App {
	return fx.New(
		fx.Provide(
			logger.New,
			config.New,
			database.NewConnectManager,
		),
		fx.Invoke(
			database.InitConnections, // init connection database ...
		),
	)
}
