package domain

import "github.com/google/uuid"

type Group struct {
	ID   int
	UUID uuid.UUID
	Name string
}
