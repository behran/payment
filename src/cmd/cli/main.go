package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"payment/internal/app"
	"payment/internal/config"
	"payment/internal/database"
)

func main() {
	containers := application.Cli()

	app := &cli.App{
		Version:              "v1.0.1",
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			{Name: "Kiselev Artem", Email: "kiselev@gmail.com"},
		},
		Commands: []*cli.Command{
			{
				Name:    "migration",
				Aliases: []string{"m"},
				Usage:   "Migration",
				Subcommands: []*cli.Command{
					{
						Name:  "postgres",
						Usage: "database",
						Subcommands: []*cli.Command{
							{
								Name: "up",
								Action: func(context *cli.Context) error {
									return database.RunPostgreMigration(config.New())
								},
							},
							{
								Name: "down",
								Action: func(context *cli.Context) error {
									return database.DownPostgreMigration(config.New())
								},
							},
						},
					},
				},
			},
		},
		Before: func(c *cli.Context) error {
			return containers.Start(c.Context)
		},
		After: func(c *cli.Context) error {
			return containers.Stop(c.Context)
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
