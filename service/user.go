package service

import (
	"log/slog"
	"fmt"

	"github.com/google/uuid"
	"github.com/mcrors/secret-santa-picker-server/domain"
)

type UserRepositoy interface {
	Get() (domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	List() ([]domain.User, error)
	Add(user domain.User) (int, error)
	Delete() (domain.User, error)
	Update(user domain.User) (domain.User, error)
}

type User struct {
	repo UserRepositoy
}

func NewUser(repo UserRepositoy) *User {
	return &User{
		repo: repo,
	}
}

func (s *User) Get() (domain.User, error) {
	return s.repo.Get()
}

func (s *User) Post(u domain.User) (string, error) {
	slog.Info("Posting user")
	err := u.Validate()
	if err != nil {
		return "", fmt.Errorf("error validating user: %w", err)
	}

	uuid := uuid.New()
	u.UUID = &uuid

	_, err = s.repo.Add(u)
	if err != nil {
		return "", fmt.Errorf("error posting from the user service: %w", err)
	}
	return u.UUID.String(), nil
}

func (s *User) Delete() (domain.User, error) {
	return s.repo.Delete()

}

func (s *User) Patch(u domain.User) (domain.User, error) {
	return s.repo.Update(u)
}

func (s *User) Exists(email string) bool {
	u, err := s.repo.GetByEmail(email)
	if err != nil {
		return false
	}
	if u == nil {
		return false
	}
	return true
}
