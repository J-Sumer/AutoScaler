package main

import (
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/api"
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/common"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	api.SetApiRoutes(e)

	e.Start(":8000")

	go common.HealthCheck()

}
