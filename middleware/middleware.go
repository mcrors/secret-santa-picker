package middleware

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mcrors/secret-santa-picker-server/iam"
)

func Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request().Clone(c.Request().Context())
			tokenString, err := iam.GetTokenFromCookie(r)
			if err != nil {
				slog.Info("could not extract token, redirecting to login:", "error", err.Error())
				return c.Redirect(http.StatusSeeOther, "/login")
			}

			claims, err := iam.VerifyToken(tokenString.Value)
			if err != nil {
				slog.Info("could not verify token, redirecting to login:", "error", err.Error())
				return c.Redirect(http.StatusSeeOther, "/login")
			}

			c.Set("username", claims.Username)

			return next(c)
		}
	}
}
