package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mcrors/secret-santa-picker-server/iam"
)

func GetIndex(c echo.Context) error {
	r := c.Request()
	w := c.Response().Writer
	isAuthenticated := iam.IsAuthenticated(r)

	if isAuthenticated {
		slog.Info("User is authenticated. Redirecting to home.")
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		slog.Info("User is not authenticated. Redirecting to login.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	return nil
}
