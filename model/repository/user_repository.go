package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zubroide/go-api-boilerplate/logger"
	"github.com/zubroide/go-api-boilerplate/model/entity"
	"github.com/zubroide/gorm-crud"
)


type UserListQueryBuilder struct {
	gorm_crud.ListQueryBuilderInterface
	*gorm_crud.BaseListQueryBuilder
}

func NewUserListQueryBuilder(db *gorm.DB, logger logger.LoggerInterface) gorm_crud.ListQueryBuilderInterface {
	base := gorm_crud.NewBaseListQueryBuilder(db, logger).(*gorm_crud.BaseListQueryBuilder)
	return &UserListQueryBuilder{BaseListQueryBuilder: base}
}

func (c UserListQueryBuilder) ListQuery(parameters gorm_crud.ListParametersInterface) (*gorm.DB, error) {
	query, err := c.BaseListQueryBuilder.ListQuery(parameters)
	params := parameters.(*UserListParameters)
	if err == nil && params.Name != "" {
		query = query.Where("name LIKE ?", params.Name + "%")
	}
	return query, err
}


type UserRepositoryInterface interface {
	gorm_crud.CrudRepositoryInterface
}

type UserListParameters struct {
	*gorm_crud.CrudListParameters
	Name string
}

type UserRepository struct {
	*gorm_crud.CrudRepository
	model entity.User
}

func NewUserRepository(db *gorm.DB, logger logger.LoggerInterface) UserRepositoryInterface {
	var model entity.User
	queryBuilder := NewUserListQueryBuilder(db, logger)
	repo := gorm_crud.NewCrudRepository(db, &model, queryBuilder, logger).(*gorm_crud.CrudRepository)
	return &UserRepository{repo, model}
}
