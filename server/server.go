// Package server is concerned with creating a http server and exposing a ListenAndServe function
package server

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mcrors/secret-santa-picker-server/config"
	"github.com/mcrors/secret-santa-picker-server/static"
	"github.com/mcrors/secret-santa-picker-server/views"
)

type Server struct {
	e               *echo.Echo
	port            string
	protectedRoutes *echo.Group
}

func NewServer(
	cfg config.HTTP,
) (*Server, error) {
	slog.Info("Creating server")

	port := fmt.Sprintf(":%s", strconv.Itoa(int(cfg.Port)))

	e := echo.New()

	s := &Server{
		e:    e,
		port: port,
	}

	err := s.setRenderers()
	if err != nil {
		return nil, err
	}

	s.configureEcho(cfg)
	s.mountHandlers()
	s.serveStaticContent()

	return s, nil
}

func (s *Server) setRenderers() error {
	t, err := views.NewTemplateManager()
	if err != nil {
		return err
	}
	s.e.Renderer = t
	return nil
}

func (s *Server) configureEcho(_ config.HTTP) {
	slog.Info("Configuring echo")
	s.e.HideBanner = true
}

func (s *Server) ListenAndServe() error {
	slog.Info("Listening on " + string(s.port))
	return s.e.Start(string(s.port))
}

func (s *Server) mountHandlers() {
	slog.Info("Mounting handlers")
}

func (s *Server) serveStaticContent() {
	content := static.Content()
	s.e.StaticFS("/static", content)
}
