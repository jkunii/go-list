package resource

import (
	"net/http"

	"github.com/jkunii/go-list/global"
	"github.com/jkunii/go-list/service"
	"github.com/labstack/echo"
)

func (tt AnnouncementResource) Get(c echo.Context) error {

	p, err := service.AdList("a")
	if err != nil {
		global.Error(err)
	}

	return c.JSON(http.StatusOK, p)
}

type AnnouncementResource struct {
}
