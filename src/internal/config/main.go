package config

import (
	"os"

	"payment/internal/tools"
)

//Config ...
type Config struct {
	App      App
	Database Database
}

//App ...
type App struct {
	Port         string
	Env          string
	Version      string
	TimeRollback int
}

//Database ...
type Database struct {
	Postgre map[int]Postgre
}

//Postgre ...
type Postgre struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

//PaymentConnect ...
const (
	PaymentConnect = iota
)

//New ...
func New() Config {
	return Config{
		App: App{
			Port:         os.Getenv("APP_PORT"),
			Version:      os.Getenv("API_VERSION"),
			TimeRollback: tools.StringToInt(os.Getenv("TIME_ROLLBACK_TX")),
		},
		Database: Database{
			Postgre: map[int]Postgre{
				PaymentConnect: {
					Host:     os.Getenv("DB_POSTGRE_HOST"),
					Port:     os.Getenv("DB_POSTGRE_PORT"),
					User:     os.Getenv("DB_POSTGRE_USER"),
					Password: os.Getenv("DB_POSTGRE_PASSWORD"),
					Database: os.Getenv("DB_POSTGRE_DATABASE"),
				},
			},
		},
	}
}
