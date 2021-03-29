package instances_test

import (
	"os"
	"testing"

	"payment/internal/config"
	"payment/internal/database/instances"
	"payment/internal/database/instances/mock"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	var m mock.RepositoryMock

	conf := config.Database{
		Postgre: map[int]config.Postgre{
			mock.ConnectMock: {
				Host:     os.Getenv("DB_POSTGRE_HOST"),
				Port:     os.Getenv("DB_POSTGRE_PORT"),
				User:     os.Getenv("DB_POSTGRE_USER"),
				Password: os.Getenv("DB_POSTGRE_PASSWORD"),
				Database: os.Getenv("DB_POSTGRE_DATABASE"),
			},
		},
	}
	instance := instances.NewPostgreSQL(conf)

	assert.NoError(t, instance.Init())

	connect, err := instance.Connect(m.ConnectName())
	if err != nil {
		assert.Error(t, err)
	}
	// check get connect ...
	assert.NotEmpty(t, connect)

	// connect not found by connect name ...
	if _, err := instance.Connect(12); err != nil {
		assert.Error(t, err)
	}
}
