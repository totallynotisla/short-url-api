package v1

import (
	"short-url/routes/v1/public"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	v1 := r.Group("/v1")

	public.Routes(v1)
}
