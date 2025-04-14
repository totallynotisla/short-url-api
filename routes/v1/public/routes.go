package public

import (
	"github.com/gin-gonic/gin"
)

var handler Handler

func Routes(r *gin.RouterGroup) {
	r.GET("/ping", handler.Ping)
}
