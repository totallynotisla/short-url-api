package main

import (
	"fmt"

	"short-url/pkg"
	"short-url/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	err := pkg.GetEnv(".env")
	if err != nil {
		panic(err)
	}
	pkg.InitConfig()

	if pkg.Config.Env == pkg.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	}
	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%d", pkg.Config.Port))
}
