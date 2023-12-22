package domain

import "github.com/google/uuid"

type SecretSanta struct {
	ID        int
	UUID      uuid.UUID
	Year      int
	Group     *Group
	Santa     *User
	Recipient *User
}
