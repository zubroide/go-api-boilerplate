package repository

import (
	"go-api-boilerplate/logger"
	"go-api-boilerplate/model/entity"
	"github.com/jinzhu/gorm"
)


type UserListQueryBuilder struct {
	ListQueryBuilderInterface
	*BaseListQueryBuilder
}

func NewUserListQueryBuilder(db *gorm.DB, logger *logger.Logger) ListQueryBuilderInterface {
	base := NewBaseListQueryBuilder(db, logger).(*BaseListQueryBuilder)
	return &UserListQueryBuilder{BaseListQueryBuilder: base}
}

func (c UserListQueryBuilder) ListQuery(parameters ListParametersInterface) *gorm.DB {
	query := c.BaseListQueryBuilder.ListQuery(parameters)
	params := parameters.(*UserListParameters)
	if params.Name != "" {
		query = query.Where("name LIKE ?", params.Name + "%")
	}
	return query
}


type UserRepositoryInterface interface {
	CrudRepositoryInterface
}

type UserListParameters struct {
	*CrudListParameters
	Name string
}

type UserRepository struct {
	*CrudRepository
	model entity.User
}

func NewUserRepository(db *gorm.DB, logger *logger.Logger) UserRepositoryInterface {
	var model entity.User
	queryBuilder := NewUserListQueryBuilder(db, logger)
	repo := NewCrudRepository(db, &model, queryBuilder, logger).(*CrudRepository)
	return &UserRepository{repo, model}
}

