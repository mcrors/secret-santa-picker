package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetHome(c echo.Context) error {
	slog.Info("Getting home page")

	username, ok := c.Get("username").(string)
	if !ok {
		slog.Info("could not get username from context")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	return c.Render(
		http.StatusOK,
		"home.html",
		struct{ Username string }{Username: username},
	)
}
