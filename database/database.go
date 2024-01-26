package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/mcrors/secret-santa-picker-server/config"
)

const (
	postgresDriver = "postgres"
	connStr        = "postgres://%s:%s@%s:%d/%s?sslmode=%s"
)

var postgresDB *sql.DB

func GetPostgresDB(cfg config.Config) (*sql.DB, error) {
	if postgresDB != nil {
		return postgresDB, nil
	}

	dataSourceName := fmt.Sprintf(connStr, cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.SSLMode)
	db, err := sql.Open(postgresDriver, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error creating postgres db: %w", err)
	}

	return db, nil
}
