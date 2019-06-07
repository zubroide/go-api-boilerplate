package repository

import (
	"go-api-boilerplate/logger"
	"go-api-boilerplate/model/entity"
	"github.com/jinzhu/gorm"
	"reflect"
)

type ListParametersInterface interface {}

type PaginationParameters struct {
	Page int
	PageSize int
}

type CrudListParameters struct {
	ListParametersInterface
	*PaginationParameters
}

const DefaultPageSize = 20


type ListQueryBuilderInterface interface {
	ListQuery(parameters ListParametersInterface) *gorm.DB
}

type BaseListQueryBuilder struct {
	db *gorm.DB
	logger *logger.Logger
	ListQueryBuilderInterface
}

func NewBaseListQueryBuilder(db *gorm.DB, logger *logger.Logger) ListQueryBuilderInterface {
	return &BaseListQueryBuilder{db: db, logger: logger}
}

func (c BaseListQueryBuilder) paginationQuery(parameters ListParametersInterface) *gorm.DB {
	query := c.db

	val := reflect.ValueOf(parameters).Elem()
	if val.Kind() != reflect.Struct {
		c.logger.Fatal("Unexpected type of parameters for paginationQuery")
		return query
	}

	var page int64
	page = 0
	pageValue := val.FieldByName("Page")
	if !pageValue.IsValid() || pageValue.Kind() != reflect.Int {
		c.logger.Fatal("Page is not specified correctly in listQuery")
	} else {
		page = pageValue.Int()
	}

	var pageSize int64
	pageSize = DefaultPageSize
	pageSizeValue := val.FieldByName("PageSize")
	if !pageSizeValue.IsValid() || pageSizeValue.Kind() != reflect.Int {
		c.logger.Fatal("PageSize is not specified in listQuery")
	} else {
		pageSize = pageSizeValue.Int()
	}

	limit := pageSize
	offset := page * pageSize
	query = query.Offset(offset).Limit(limit)

	return query
}

func (c BaseListQueryBuilder) ListQuery(parameters ListParametersInterface) *gorm.DB {
	return c.paginationQuery(parameters)
}


type CrudRepositoryInterface interface {
	BaseRepositoryInterface
	GetModel() (entity.InterfaceEntity)
	Find(id uint) (entity.InterfaceEntity, error)
	List(parameters ListParametersInterface) (entity.InterfaceEntity, error)
	Create(item entity.InterfaceEntity) entity.InterfaceEntity
	Update(item entity.InterfaceEntity) entity.InterfaceEntity
	Delete(id uint) error
}

type CrudRepository struct {
	CrudRepositoryInterface
	*BaseRepository
	model entity.InterfaceEntity // Dynamic typing
	listQueryBuilder ListQueryBuilderInterface
}

func NewCrudRepository(db *gorm.DB, model entity.InterfaceEntity, listQueryBuilder ListQueryBuilderInterface, logger *logger.Logger) CrudRepositoryInterface {
	repo := NewBaseRepository(db, logger).(*BaseRepository)
	return &CrudRepository{
		BaseRepository: repo,
		model: model,
		listQueryBuilder: listQueryBuilder,
	}
}

func (c CrudRepository) GetModel() (entity.InterfaceEntity) {
	return c.model
}

func (c CrudRepository) Find(id uint) (entity.InterfaceEntity, error) {
	item := reflect.New(reflect.TypeOf(c.GetModel()).Elem()).Interface()
	err := c.db.First(item, id).Error
	return item, err
}

func (c CrudRepository) List(parameters ListParametersInterface) (entity.InterfaceEntity, error) {
	items := reflect.New(reflect.SliceOf(reflect.TypeOf(c.GetModel()).Elem())).Interface()
	query := c.listQueryBuilder.ListQuery(parameters)
	err := query.Find(items).Error
	return items, err
}

func (c CrudRepository) Create(item entity.InterfaceEntity) entity.InterfaceEntity {
	c.db.Create(item)
	return item
}

func (c CrudRepository) Update(item entity.InterfaceEntity) entity.InterfaceEntity {
	c.db.Save(item)
	return item
}

func (c CrudRepository) Delete(id uint) error {
	item, err := c.Find(id)
	if err != nil {
		return err
	}
	c.db.Delete(item)
	return nil
}
