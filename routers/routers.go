package routers

import (
	"github.com/jkunii/go-list/global"
	"github.com/jkunii/go-list/resource"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type ApplicationRouter struct {
	Announcement *resource.AnnouncementResource `inject:""`
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

	g.GET("/list", r.Announcement.Get)

}
