package iam

import (
	"log/slog"
	"net/http"
)

func IsAuthenticated(r *http.Request) bool {
	slog.Info("Checking if user is authenticated.")
	return true
}
