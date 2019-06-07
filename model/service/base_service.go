package service

import (
	"go-api-boilerplate/logger"
	"go-api-boilerplate/model/repository"
)

type BaseServiceInterface interface {
}

type BaseService struct {
	BaseServiceInterface
	repository repository.BaseRepositoryInterface
	logger *logger.Logger
}

func NewBaseService(repository repository.BaseRepositoryInterface, logger *logger.Logger) *BaseService {
	return &BaseService{repository: repository, logger: logger}
}
