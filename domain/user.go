package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID        int
	UUID      uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
}
