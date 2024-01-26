package postgres_test

import (
	"testing"
	"github.com/mcrors/secret-santa-picker-server/repository/postgres"
)

func TestAddUserToRepository(t *testing.T) {
	db := postgres.NewDB()
	repo := postgres.NewUserRepository(db)
	user := postgres.User{
		Email: "

