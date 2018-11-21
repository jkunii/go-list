package routers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jkunii/go-list/global"
	"github.com/jkunii/go-list/resource"
)

type ApplicationRouter struct {
	Property         *resource.PropertyResource         `inject:""`

}

func (r ApplicationRouter) Init(e *echo.Echo) {

	// Swagger
	e.Static("/", "public/swagger/")

	g := e.Group("/api")

	// g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) bool {
	// g.Use(middleware.BasicAuth(func(username, password string) bool {
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == global.Cfg.UserName && password == global.Cfg.Secret {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/list", r.Property.Get)

}
