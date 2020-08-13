package service

import (
	"github.com/zubroide/go-api-boilerplate/logger"
	"github.com/zubroide/go-api-boilerplate/model/repository"
	"github.com/zubroide/gorm-crud"
)

type UserServiceInterface interface {
	gorm_crud.CrudServiceInterface
}

type UserService struct {
	*gorm_crud.CrudService
	repository repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface, logger logger.LoggerInterface) UserServiceInterface {
	crudService := gorm_crud.NewCrudService(repository, logger).(*gorm_crud.CrudService)
	service := &UserService{crudService, repository}
	return service
}
