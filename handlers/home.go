package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetHome(c echo.Context) error {
	slog.Info("Getting home page")

	if err := c.Render(http.StatusOK, "home.html", nil); err != nil {
		slog.Error("Error rendering home page: " + err.Error())
		return err
	}
	return nil
}
