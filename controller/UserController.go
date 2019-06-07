package controller

import (
	"go-api-boilerplate/logger"
	"go-api-boilerplate/model/repository"
	"go-api-boilerplate/model/service"
	"github.com/gin-gonic/gin"
)


type UserListParametersHydrator struct {
	*BaseParametersHydrator
}

func NewUserListParametersHydrator(logger *logger.Logger) ParametersHydratorInterface {
	base := NewBaseParametersHydrator(logger).(*BaseParametersHydrator)
	return &UserListParametersHydrator{BaseParametersHydrator: base}
}

func (c UserListParametersHydrator) Hydrate(context *gin.Context) (repository.ListParametersInterface, error) {
	crudParams, _ := c.BaseParametersHydrator.Hydrate(context)
	parameters := &repository.UserListParameters{
		CrudListParameters: crudParams.(*repository.CrudListParameters),
	}
	if err := context.ShouldBindQuery(parameters); err != nil {
		return crudParams, err
	}

	return parameters, nil
}


type UserController struct {
	*CrudController
	service service.UserServiceInterface
}

func NewUserController(service service.UserServiceInterface, logger *logger.Logger) *UserController {
	parametersHydrator := NewUserListParametersHydrator(logger)
	controller := NewCrudController(service, parametersHydrator, logger)
	return &UserController{CrudController: controller, service: service}
}
