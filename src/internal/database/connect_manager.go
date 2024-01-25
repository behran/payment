package database

import (
	"context"

	"go.uber.org/fx"
	"payment/internal/config"
	"payment/internal/database/instances"

	"github.com/jmoiron/sqlx"

	"golang.org/x/sync/errgroup"
)

// ConnectManager ...
type ConnectManager struct {
	postgre *instances.PostgreSQL
}

// NewConnectManager ...
func NewConnectManager(config config.Config) *ConnectManager {
	return &ConnectManager{
		postgre: instances.NewPostgreSQL(config.Database),
	}
}

// InitConnections ...
func InitConnections(lc fx.Lifecycle, manager *ConnectManager) error {
	var eg errgroup.Group

	for _, instance := range []func() error{
		manager.postgre.Init,
	} {
		eg.Go(instance)
	}
	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return manager.DisconnectPostgreSQL()
		},
	})
	return eg.Wait()
}

// ConnectPostgreSQL ...
func (c ConnectManager) ConnectPostgreSQL(client IClient) (*sqlx.DB, error) {
	return c.postgre.Connect(client.ConnectName())
}

func (c ConnectManager) DisconnectPostgreSQL() error {
	return c.postgre.Disconnect()
}
