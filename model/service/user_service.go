package service

import (
	"github.com/zubroide/go-api-boilerplate/logger"
	"github.com/zubroide/go-api-boilerplate/model/entity"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetUsers(string) ([]*entity.User, error)
}

type UserService struct {
	db     *gorm.DB
	logger logger.LoggerInterface
}

func NewUserService(db *gorm.DB, logger logger.LoggerInterface) UserServiceInterface {
	service := &UserService{db, logger}
	return service
}

func (s *UserService) GetUsers(name string) ([]*entity.User, error) {
	var items []*entity.User
	res := s.db.
		Where("name ilike ?||'%'", name).
		Order("name").
		Find(&items)
	if res.Error != nil {
		return nil, res.Error
	}
	return items, nil
}
