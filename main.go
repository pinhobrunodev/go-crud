package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pinhobrunodev/go-crud/src/configuration/logger"
	"github.com/pinhobrunodev/go-crud/src/controller/routes"
	"github.com/pinhobrunodev/go-crud/utils"
)

func main() {
	logger.Info(utils.INIT_APPLICATION)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
