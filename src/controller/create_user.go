package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinhobrunodev/go-crud/src/configuration/logger"
	"github.com/pinhobrunodev/go-crud/src/configuration/validation"
	"github.com/pinhobrunodev/go-crud/src/controller/model/request"
	"github.com/pinhobrunodev/go-crud/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		rest_err := validation.ValidateUserError(err)
		logger.Error("Error trying to validade user info", err,
			zap.String("journey", "createUser"))
		c.JSON(rest_err.Code, rest_err)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully!", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
