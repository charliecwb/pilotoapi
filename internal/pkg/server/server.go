package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *HttpServerSvc) Serve(app *echo.Echo, url string) error {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	//app.Use(func(c *echo.Echo) error {
	//	return c.SendStatus(fiber.StatusNotFound)
	//})

	return app.Start(url)
}
