// Package domain contains the domain models
package domain

import (
	"time"

	"github.com/google/uuid"
)

type ID = uuid.UUID

type Group struct {
	ID        ID
	Name      string
	CreatedAt time.Time
}

type Member struct {
	ID      ID
	GroupID ID
	Name    string
	Email   *string
}

type Event struct {
	ID      ID
	GroupID ID
	Year    int
	// later: status, drawDate, rules snapshot, etc.
}

type Assignment struct {
	ID         ID
	EventID    ID
	GiverID    ID
	ReceiverID ID
}
