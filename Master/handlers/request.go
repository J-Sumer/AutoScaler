package handlers

import (
	"net/http"

	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/common"
	"github.com/labstack/echo"
)

func ProxyRequest(c echo.Context) error {
	common.Lb(c.Response().Writer, c.Request())
	return c.JSON(http.StatusOK, "Message: Ok")
}
