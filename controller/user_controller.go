package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zubroide/gin-crud"
	"github.com/zubroide/go-api-boilerplate/logger"
	"github.com/zubroide/go-api-boilerplate/model/repository"
	"github.com/zubroide/go-api-boilerplate/model/service"
	"github.com/zubroide/gorm-crud"
)


type UserListParametersHydrator struct {
	*gin_crud.BaseParametersHydrator
}

func NewUserListParametersHydrator(logger logger.LoggerInterface) gin_crud.ParametersHydratorInterface {
	base := gin_crud.NewBaseParametersHydrator(logger).(*gin_crud.BaseParametersHydrator)
	return &UserListParametersHydrator{BaseParametersHydrator: base}
}

func (c UserListParametersHydrator) Hydrate(context *gin.Context) (gorm_crud.ListParametersInterface, error) {
	crudParams, _ := c.BaseParametersHydrator.Hydrate(context)
	parameters := &repository.UserListParameters{
		CrudListParameters: crudParams.(*gorm_crud.CrudListParameters),
	}
	if err := context.ShouldBindQuery(parameters); err != nil {
		return crudParams, err
	}

	return parameters, nil
}


type UserController struct {
	*gin_crud.CrudController
	service service.UserServiceInterface
}

func NewUserController(service service.UserServiceInterface, logger logger.LoggerInterface) *UserController {
	parametersHydrator := NewUserListParametersHydrator(logger)
	controller := gin_crud.NewCrudController(service, parametersHydrator, logger)
	return &UserController{CrudController: controller, service: service}
}
