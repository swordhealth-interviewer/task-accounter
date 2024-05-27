package restapi

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (s *RestServer) SetUpRoutes(config echojwt.Config) {
	v1 := s.router.Group("/v1")
	s.LoginRoute(v1)

	v2 := s.router.Group("/v2")
	v2.Use(echojwt.WithConfig(config))
	s.PingRoute(v2)
	s.TaskRoutes(v2)
}

func (s *RestServer) LoginRoute(v1 *echo.Group) {
	v1.POST("/login", func(c echo.Context) error {
		return s.appHandler.loginHandler.Handle(c)
	})
}

func (s *RestServer) PingRoute(v2 *echo.Group) {
	v2.GET("/ping", func(c echo.Context) error {
		return s.appHandler.pingHandler.Handle(c)
	})
}

func (s *RestServer) TaskRoutes(v2 *echo.Group) {
	v2.POST("/task", func(c echo.Context) error {
		return s.appHandler.taskCreateHandler.Handle(c)
	})
}
