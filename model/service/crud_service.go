package service

import (
	"go-api-boilerplate/logger"
	"go-api-boilerplate/model/entity"
	"go-api-boilerplate/model/repository"
)

type CrudServiceInterface interface {
	BaseServiceInterface
	GetModel() (entity.InterfaceEntity)
	GetItem(id uint) (entity.InterfaceEntity, error)
	GetList(parameters repository.ListParametersInterface) (entity.InterfaceEntity, error)
	Create(item entity.InterfaceEntity) entity.InterfaceEntity
	Update(item entity.InterfaceEntity) entity.InterfaceEntity
	Delete(id uint) error
}

type CrudService struct {
	*BaseService
	repository repository.CrudRepositoryInterface
}

func NewCrudService(repository repository.CrudRepositoryInterface, logger *logger.Logger) CrudServiceInterface {
	service := NewBaseService(repository, logger)
	return &CrudService{service, repository}
}

func (c CrudService) GetModel() (entity.InterfaceEntity) {
	return c.repository.GetModel()
}

func (c CrudService) GetItem(id uint) (entity.InterfaceEntity, error) {
	return c.repository.Find(id)
}

func (c CrudService) GetList(parameters repository.ListParametersInterface) (entity.InterfaceEntity, error) {
	return c.repository.List(parameters)
}

func (c CrudService) Create(item entity.InterfaceEntity) entity.InterfaceEntity {
	return c.repository.Create(item)
}

func (c CrudService) Update(item entity.InterfaceEntity) entity.InterfaceEntity {
	return c.repository.Update(item)
}

func (c CrudService) Delete(id uint) error {
	return c.repository.Delete(id)
}
