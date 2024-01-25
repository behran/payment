package instances

import (
	"fmt"
	"time"

	"payment/internal/config"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/jmoiron/sqlx"
)

// PostgreSQL ...
type PostgreSQL struct {
	config config.Database
	pool   map[int]*sqlx.DB
}

// NewPostgreSQL ...
func NewPostgreSQL(config config.Database) *PostgreSQL {
	return &PostgreSQL{
		config: config,
		pool:   make(map[int]*sqlx.DB),
	}
}

// Init ...
func (m *PostgreSQL) Init() error {
	for k := range m.config.Postgre {
		connect, err := m.create(m.config.Postgre[k])
		if err != nil {
			return err
		}
		m.append(k, connect)
	}
	return nil
}

func (m PostgreSQL) create(config config.Postgre) (*sqlx.DB, error) {
	schema := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	connect, err := sqlx.Connect("pgx", schema)
	if err != nil {
		return nil, err
	}
	connect.SetMaxOpenConns(25)
	connect.SetMaxIdleConns(25)
	connect.SetConnMaxLifetime(5 * time.Minute)

	return connect, nil
}

// Connect ...
func (m PostgreSQL) Connect(name int) (*sqlx.DB, error) {
	connect, exist := m.pool[name]
	if !exist {
		return nil, fmt.Errorf("connect postgresql `%d` not found", name)
	}
	return connect, nil
}

func (m PostgreSQL) Disconnect() error {
	for k := range m.pool {
		if err := m.pool[k].Close(); err != nil {
			return err
		}
	}
	return nil
}

func (m *PostgreSQL) append(key int, connect *sqlx.DB) { m.pool[key] = connect }
