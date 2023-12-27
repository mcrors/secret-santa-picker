package domain

import "github.com/google/uuid"

type Member struct {
	ID    int
	UUID  uuid.UUID
	User  *User
	Group *Group
}
