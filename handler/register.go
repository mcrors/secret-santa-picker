package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
	slog.Info("PostRegister")
	var req UserPostRequestData
	err := c.Bind(&req)
	if err != nil {
		slog.Error("error binding request data: " + err.Error())
		return err
	}

	_, err = r.userService.Post(*req.ToUser())
	if err != nil {
		slog.Error("error posting user: " + err.Error())
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}
