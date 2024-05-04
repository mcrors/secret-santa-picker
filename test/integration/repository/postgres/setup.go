package postgres_test

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/mcrors/secret-santa-picker-server/config"
)

func testConfig() config.Config {
	return config.Config{
		App: config.App{
			LogLevel: "DEBUG",
		},
		Http: config.HTTP{
			Port: 8080,
			Host: "localhost",
		},
		DB: config.Database{
			Host:     "localhost",
			Port:     5432,
			Username: "secret_santa_user",
			Password: "secret_santa_password",
			Name:     "secret_santa_db",
			SSLMode:  "disable",
		},
	}
}

func dbMigration(db *sql.DB) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	return migrate.NewWithDatabaseInstance("file://../../../../database/migration", "postgres", driver)
}
