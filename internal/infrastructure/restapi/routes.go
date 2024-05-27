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

	task := v2.Group("/task")
	s.TaskRoutes(task)
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

func (s *RestServer) TaskRoutes(task *echo.Group) {
	task.POST("/", func(c echo.Context) error {
		return s.appHandler.taskCreateHandler.Handle(c)
	})

	task.GET("/:id", func(c echo.Context) error {
		return s.appHandler.taskReadHandler.Handle(c)
	})

	task.GET("/", func(c echo.Context) error {
		return s.appHandler.taskReadAllHandler.Handle(c)
	})

	task.PUT("/:id", func(c echo.Context) error {
		return s.appHandler.taskUpdateHandler.Handle(c)
	})

	task.DELETE("/:id", func(c echo.Context) error {
		return s.appHandler.taskDeleteHandler.Handle(c)
	})
}
