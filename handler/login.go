package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Lo

type Login struct {
	userService UserService
}

func NewLogin(u UserService) *Login {
	return &Login{
		userService: u,
	}
}

func (l *Login) GetLogin(c echo.Context) error {
	slog.Info("GetLogin")
	return c.Render(http.StatusOK, "login.html", nil)
}

func (l *Login) PostLogin(c echo.Context) error {
	slog.Info("PostLogin")
	var req LoginPostRequestData
	err := c.Bind(&req)
	if err != nil {
		slog.Error("error binding post login request data: " + err.Error())
		return err
	}

}
