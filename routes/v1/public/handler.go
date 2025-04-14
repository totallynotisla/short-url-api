package public

import (
	"net/http"
	"short-url/pkg"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, pkg.ApiResponse{
		Success: true,
		Message: "Pong!",
	})
}
