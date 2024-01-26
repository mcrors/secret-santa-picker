package database

import (
	"database/sql"
	"fmt"

	"github.com/mcrors/secret-santa-picker-server/config"
)

type PostgresDB struct {
	*sql.DB
}

var postgresDB *PostgresDB

func GetPostgresDB(cfg config.Config) (*PostgresDB, error) {
	if postgresDB != nil {
		return postgresDB, nil
	}

	db, err := sql.Open("postgres", "not implemented")
	if err != nil {
		return nil, fmt.Errorf("error creating postgres db: %w", err)
	}

	return &PostgresDB{
		DB: db,
	}, nil
}
