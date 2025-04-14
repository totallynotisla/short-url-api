package routes

import (
	"net/http"
	"short-url/pkg"

	"github.com/gin-gonic/gin"
)

type Middleware struct{}

func (Middleware) SetUp(r *gin.Engine) {
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
}
