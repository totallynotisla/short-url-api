package routes

import (
	"short-url/pkg"
	v1 "short-url/routes/v1"

	"github.com/gin-gonic/gin"
)

var rootMiddleware Middleware

func SetupRouter() *gin.Engine {
	r := gin.New()
	if pkg.Config.Env == pkg.EnvProduction {
		r.SetTrustedProxies([]string{"127.0.0.1"})
	}

	rootMiddleware.SetUp(r)

	v1.Routes(r)

	return r
}
