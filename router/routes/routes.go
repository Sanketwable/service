package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

type Route struct {
	Uri     string
	Handler echo.HandlerFunc
}

func SetUpRoutes(e *echo.Echo) *echo.Echo {
	R := Load()
	// for _, route := range Load() {
	// 	e.POST(route.Uri, route.Handler)
	// }
	e.POST(R[0].Uri, R[0].Handler)
	e.GET(R[1].Uri, R[1].Handler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}

func Load() []Route {
	routes := covid_cases
	routes = append(routes, fetch_cases...)
	return routes
}
