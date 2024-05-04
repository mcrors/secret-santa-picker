package postgres_test

import (
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/mcrors/secret-santa-picker-server/database"
	"github.com/mcrors/secret-santa-picker-server/domain"
	pg_repo "github.com/mcrors/secret-santa-picker-server/repository/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/ory/dockertest/v3"
	dc "github.com/ory/dockertest/v3/docker"
)

func TestMain(m *testing.M) {
	cfg := testConfig()
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not construct a pool: %v", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("could not connect to docker: %v", err)
	}

	resource, err := pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: "postgres",
			Tag:        "16",
			Name:       "postgres_test",
			PortBindings: map[dc.Port][]dc.PortBinding{
				"5432/tcp": {{HostIP: "", HostPort: "5432"}},
			},
			Env: []string{
				"POSTGRES_USER=" + cfg.DB.Username,
				"POSTGRES_PASSWORD=" + cfg.DB.Password,
				"POSTGRES_DB=" + cfg.DB.Name,
			},
		},
		func(config *dc.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = dc.RestartPolicy{Name: "no"}
		},
	)

	if err != nil {
		log.Fatalf("could not start resource: %v", err)
	}

	if err = pool.Retry(func() error {
		db, err := database.GetPostgresDB(cfg.DB)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	exitCode := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not purge resource: %v", err)
	}

	os.Exit(exitCode)
}

func TestAddUserToRepository(t *testing.T) {
	cfg := testConfig()

	db, err := database.GetPostgresDB(cfg.DB)
	if err != nil {
		t.Fatalf("error creating postgres db: %v", err)
	}
	defer db.Close()

	m, err := dbMigration(db)
	if err != nil {
		t.Fatalf("error creating migration: %v", err)
	}
	defer m.Down()
	if err = m.Up(); err != nil {
		t.Fatalf("error running migration: %v", err)
	}

	repo := pg_repo.NewUser(db)
	uuid := uuid.New()
	user := domain.User{
		UUID:         &uuid,
		FirstName:    "Test",
		LastName:     "User",
		Email:        "test.user@email.com",
		PasswordHash: "password",
	}
	id, err := repo.Add(user)
	if err != nil {
		t.Fatalf("error adding user to repository: %v", err)
	}
	if id != 1 {
		t.Fatalf("expected id to be 1, got %d", id)
	}
}
