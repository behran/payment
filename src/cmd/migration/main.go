package main

import (
	"fmt"
	"log"

	"payment/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	conf := config.New().Database.Postgre[config.PaymentConnect]
	log.Printf("%+v", conf)
	schema := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)
	m, err := migrate.New(
		"file://migrations",
		schema)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
