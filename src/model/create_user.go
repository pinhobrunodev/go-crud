package model

import (
	"github.com/pinhobrunodev/go-crud/src/configuration/logger"
	"github.com/pinhobrunodev/go-crud/src/configuration/rest_err"
	"github.com/pinhobrunodev/go-crud/utils"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info(utils.INIT_CREATE_USER_MODEL, zap.String(utils.JOURNEY, utils.CREATE_USER))
	ud.EncryptPassword()
	return nil
}
