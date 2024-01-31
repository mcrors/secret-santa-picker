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

func GetPostgresDB(dbConfig config.Database) (*sql.DB, error) {
	if postgresDB != nil {
		return postgresDB, nil
	}

	dataSourceName := fmt.Sprintf(
		connStr,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.SSLMode,
	)
	db, err := sql.Open(postgresDriver, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error creating postgres db: %w", err)
	}

	return db, nil
}
