package routes

import (
	"net/http"
	// "strconv"

	echo "github.com/labstack/echo/v4"
	"github.com/J-Sumer/AutoScaler/velvet/docker"
)

func RunningContainersCountRoute(c echo.Context) error {
	return c.String(http.StatusOK, RunningContainersCount())
}

func GetContainerMapping(c echo.Context) error {
	return c.String(http.StatusOK, ContainerMappingMap())
}

func GetContainerList(c echo.Context) error {
	return c.JSON(http.StatusOK, docker.RunningContainers())
}

func CreateContainerRoute(c echo.Context) error {
  	// User ID from path `users/:id`
  	// port := c.Param("port")
  	name := c.Param("name")
	return c.JSON(http.StatusOK, CreateContainer(name))
}

func DeleteContainerRoute(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, DeleteContainer(id))
}

func HelloWorldRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

// func AddMetricEntryRoute(c echo.Context) error {
// 	cpu, _ := strconv.Atoi(c.Param("cpu"))
// 	memory, _ := strconv.Atoi(c.Param("memory"))
// 	return c.String(http.StatusOK, AddMetricEntry(cpu, memory))
// }

func MetricsRoute(c echo.Context) error {
	return c.String(http.StatusOK, GetCPUMetric())
}