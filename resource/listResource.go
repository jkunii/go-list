package resource

import (
	"net/http"
	"strconv"

	"github.com/jkunii/go-list/global"
	"github.com/jkunii/go-list/service"
	"github.com/labstack/echo"
)

func (tt AnnouncementResource) Get(c echo.Context) error {

	groupType := c.QueryParam("type")
	offSet, _ := strconv.Atoi(c.QueryParam("offSet"))
	limitPerPage, _ := strconv.Atoi(c.QueryParam("limit"))
	orderByPrice := c.QueryParam("orderByPrice")
	p, err := service.AdList(groupType, orderByPrice, offSet, limitPerPage)
	if err != nil {
		global.Error(err)
	}

	return c.JSON(http.StatusOK, p)
}

type AnnouncementResource struct {
}
