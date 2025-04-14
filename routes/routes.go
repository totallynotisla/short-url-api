package route

import (
	"net/http"
	"short-url/pkg"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, pkg.ApiResponse{
			Success: true,
			Message: "Pong!",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, pkg.ApiResponse{
			Success: false,
			Message: "Resource not found",
		})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, pkg.ApiResponse{
			Success: false,
			Message: "Method not allowed",
		})
	})

	return r
}
