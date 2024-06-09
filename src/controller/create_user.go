package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pinhobrunodev/go-crud/src/configuration/rest_err"
	"github.com/pinhobrunodev/go-crud/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		rest_err := rest_err.NewBadRequestError(
			fmt.Sprintf("The are some incorrect fields, erros=%s", err.Error()))
		c.JSON(rest_err.Code, rest_err)
		return

	}

	fmt.Println(userRequest)
}
