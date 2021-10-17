package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs"
)


type Route struct {
	Uri     string
	Handler echo.HandlerFunc
}

func SetUpRoutes(e *echo.Echo) *echo.Echo {
	for _, route := range Load() {
		e.POST(route.Uri, route.Handler)
	}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}

func Load() []Route {
	routes := covid_cases
	return routes
}