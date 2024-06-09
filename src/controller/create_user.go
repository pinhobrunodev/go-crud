package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pinhobrunodev/go-crud/src/configuration/rest_err"
	"github.com/pinhobrunodev/go-crud/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		rest_err := rest_err.NewBadRequestError("Some fields are incorrect")
		log.Printf("Error: %s", err)
		c.JSON(rest_err.Code, rest_err)
		return

	}

	fmt.Println(userRequest)
}
