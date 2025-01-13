package handler

import "github.com/mcrors/secret-santa-picker-server/domain"

type UserPostRequestData struct {
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

func (r *UserPostRequestData) ToUser() *domain.User {
	user := &domain.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
	}
	user.HashPassword(r.Password)
	return user
}

type UserGetRequestData struct {
}

type UserDeleteRequestData struct {
}

type UserPatchRequestData struct {
}

type LoginPostRequestData struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
