package http

import (
	"fmt"

	handlers "github.com/dezh-tech/panda/deliveries/http/handlers/domain"
	service "github.com/dezh-tech/panda/services/domain"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Router   *echo.Echo
	config   Config
	handlers Handlers
}

func New(config Config, userSvc service.Domain) Server {
	return Server{
		Router: echo.New(),
		config: config,

		handlers: Handlers{
			domain: handlers.NewDomainService(userSvc),
		},
	}
}

func (s Server) Start() error {
	s.handlers.Start(s.Router)

	address := fmt.Sprintf(":%d", s.config.Port)
	if err := s.Router.Start(address); err != nil {
		return err
	}

	return nil
}

func (s Server) Stop() error {
	if err := s.Router.Close(); err != nil {
		return err
	}

	return nil
}
