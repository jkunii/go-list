package resource

import (
	"net/http"

	"github.com/labstack/echo"
)

func (tt ListResource) Get(c echo.Context) error {

	return c.JSON(http.StatusOK, "works")
}

type ListResource struct {
}

type (
	list struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
