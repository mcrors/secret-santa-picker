package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mcrors/secret-santa-picker-server/domain"
)

type UserService interface {
	Post(u domain.User) (string, error)
}

type Register struct {
	userService UserService
}

func NewRegister(userService UserService) *Register {
	return &Register{
		userService: userService,
	}
}

func (r *Register) GetRegister(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}

func (r *Register) PostRegister(c echo.Context) error {
	var req UserPostRequestData
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	_, err = r.userService.Post(*req.ToUser())
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}
