package postgres

import (
	"database/sql"
	"fmt"

	"github.com/mcrors/secret-santa-picker-server/domain"
)

type User struct {
	db        *sql.DB
	tableName string
	schema    string
}

func NewUser(db *sql.DB) *User {
	return &User{
		db:        db,
		tableName: "users",
		schema:    "secret_santa",
	}
}

func (u *User) Get() (domain.User, error) {
	return domain.User{}, nil
}

func (u *User) List() ([]domain.User, error) {
	return []domain.User{}, nil
}

func (u *User) Exists(email string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM ` + u.schema + `.` + u.tableName + `
			WHERE email = $1
		)
	`
	var exists bool
	err := u.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if user exists in repository: %w", err)
	}

	return exists, nil
}

func (u *User) Add(user domain.User) (int, error) {
	query := `
		INSERT INTO ` + u.schema + `.` + u.tableName + `
		(uuid, first_name, last_name, email, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var id int
	err := u.db.QueryRow(
		query,
		user.UUID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error adding user to repository: %w", err)
	}

	return id, nil
}

func (u *User) Delete() (domain.User, error) {
	return domain.User{}, nil
}

func (u *User) Update(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}
