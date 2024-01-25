package database

import (
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"payment/internal/config"
	"payment/internal/database/instances"
)

// RunPostgreMigration ...
func RunPostgreMigration(conf config.Config) error {
	instance := instances.NewPostgreSQL(conf.Database)

	if err := instance.Init(); err != nil {
		return err
	}
	connect, err := instance.Connect(config.PaymentConnect)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(connect.DB, &postgres.Config{
		DatabaseName:     conf.Database.Postgre[config.PaymentConnect].Database,
		StatementTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		conf.Database.Postgre[config.PaymentConnect].Database,
		driver,
	)
	if err != nil {
		return err
	}
	return m.Up()
}

// DownPostgreMigration ...
func DownPostgreMigration(conf config.Config) error {
	instance := instances.NewPostgreSQL(conf.Database)

	if err := instance.Init(); err != nil {
		return err
	}
	connect, err := instance.Connect(config.PaymentConnect)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(connect.DB, &postgres.Config{
		DatabaseName:     conf.Database.Postgre[config.PaymentConnect].Database,
		StatementTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		conf.Database.Postgre[config.PaymentConnect].Database,
		driver,
	)
	if err != nil {
		return err
	}
	return m.Down()
}
