package main

import (
	"log"

	"github.com/Ashu23042000/coffee-supply-chain/backend/configuration"
	"github.com/Ashu23042000/coffee-supply-chain/backend/controller"
	"github.com/labstack/echo/v4"
)

func main() {

	configuration.LoadEnvVariables()

	e := echo.New()
	g := e.Group("/api")

	g.GET("/assets", controller.Get)
	g.GET("/asset/:id", controller.GetById)
	g.POST("/create-asset", controller.Create)

	err := e.Start(":8080")
	if err != nil {
		log.Fatalf("error while starting server:%v", err)
	}
}
