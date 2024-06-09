package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinhobrunodev/go-crud/src/configuration/logger"
	"github.com/pinhobrunodev/go-crud/src/configuration/validation"
	"github.com/pinhobrunodev/go-crud/src/controller/model/request"
	"github.com/pinhobrunodev/go-crud/src/model"
	"github.com/pinhobrunodev/go-crud/utils"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info(utils.INIT_CREATE_USER_CONTROLLER, zap.String(utils.JOURNEY, utils.CREATE_USER))
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		rest_err := validation.ValidateUserError(err)
		logger.Error(utils.USER_VALIDATION_ERROR, err,
			zap.String(utils.JOURNEY, utils.CREATE_USER))
		c.JSON(rest_err.Code, rest_err)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(utils.USER_CREATED_SUCCESSFULLY, zap.String(utils.JOURNEY, utils.CREATE_USER))

	c.String(http.StatusCreated, "Success!")
}
