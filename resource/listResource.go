package resource

import (
	"net/http"
	"strconv"

	"github.com/jkunii/go-list/global"
	"github.com/jkunii/go-list/service"
	"github.com/labstack/echo"
)

func (tt AnnouncementResource) Get(c echo.Context) error {

	venture := c.QueryParam("venture")
	groupType := c.QueryParam("type")
	offSet, _ := strconv.Atoi(c.QueryParam("offSet"))
	limitPerPage, _ := strconv.Atoi(c.QueryParam("limit"))
	orderByPrice := c.QueryParam("orderByPrice")
	p, err := service.AdList(venture, groupType, orderByPrice, offSet, limitPerPage)
	if err != nil {
		global.Error(err)
	}

	return c.JSON(http.StatusOK, p)
}

type AnnouncementResource struct {
}
