package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"short-url/pkg"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

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

func main() {
	err := pkg.GetEnv(".env")
	if err != nil {
		panic(err)
	}
	pkg.InitConfig()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%d", pkg.Config.Port))
}
