package repository

import (
	"go-api-boilerplate/logger"
	"github.com/jinzhu/gorm"
)

type BaseRepositoryInterface interface {
}

type BaseRepository struct {
	BaseRepositoryInterface
	db *gorm.DB
	logger *logger.Logger
}

func NewBaseRepository(db *gorm.DB, logger *logger.Logger) BaseRepositoryInterface {
	return &BaseRepository{db: db, logger: logger}
}
