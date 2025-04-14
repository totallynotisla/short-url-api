package main

import (
	"fmt"

	"short-url/pkg"
	route "short-url/routes"
)

func main() {
	err := pkg.GetEnv(".env")
	if err != nil {
		panic(err)
	}
	pkg.InitConfig()

	r := route.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%d", pkg.Config.Port))
}
