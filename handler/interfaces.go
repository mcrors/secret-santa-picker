package handler

import (
	"github.com/mcrors/secret-santa-picker-server/domain"
)

type UserService interface {
	Post(u domain.User) (string, error)
}
