package main

import (
	"log"

	"github.com/pinhobrunodev/go-crud/controller/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.InitRoutes()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
