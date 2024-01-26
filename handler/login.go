package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetLogin(c echo.Context) error {
	slog.Info("GetLogin")
	return c.Render(http.StatusOK, "login.html", nil)
}
