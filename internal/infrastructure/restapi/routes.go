package restapi

import "github.com/labstack/echo/v4"

func (s *RestServer) SetUpRoutes() {
	v1 := s.router.Group("/v1")
	s.PingRoute(v1)
}

func (s *RestServer) PingRoute(v1 *echo.Group) {
	v1.GET("/ping", func(c echo.Context) error {
		return s.appHandler.pingHandler.Handle(c)
	})
}
