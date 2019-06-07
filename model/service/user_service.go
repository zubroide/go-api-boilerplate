package service

import (
	"go-api-boilerplate/logger"
	"go-api-boilerplate/model/repository"
)

type UserServiceInterface interface {
	CrudServiceInterface
}

type UserService struct {
	*CrudService
	repository repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface, logger *logger.Logger) UserServiceInterface {
	crudService := NewCrudService(repository, logger).(*CrudService)
	service := &UserService{crudService, repository}
	return service
}
