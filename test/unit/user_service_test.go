package unit

import (
	"github.com/zubroide/go-api-boilerplate/dic"
	"github.com/zubroide/go-api-boilerplate/model/entity"
	"github.com/zubroide/go-api-boilerplate/model/repository"
	"github.com/zubroide/go-api-boilerplate/model/service"
	"github.com/zubroide/gorm-crud"
	"testing"
)

type UserRepositoryMock struct {
	gorm_crud.CrudRepositoryInterface
	model gorm_crud.InterfaceEntity // Dynamic typing
}

func NewUserRepositoryMock() repository.UserRepositoryInterface {
	var model entity.User
	repo := &UserRepositoryMock{model: model}
	return repo
}

func (c UserRepositoryMock) Find(id uint) (gorm_crud.InterfaceEntity, error) {
	item := &entity.User{UserFields: &entity.UserFields{Name: "test user"}}
	item.ID = 2
	return item, nil
}

func TestGetItem(t *testing.T) {
	// It is not good way, because we cannot override provided type with another.
	// So we must specify here all services
	dic.InitBuilder()
	dic.Builder.Set(
		dic.UserRepository,
		NewUserRepositoryMock(),
	)
	dic.Container = dic.Builder.Build()

	userService := dic.Container.Get(dic.UserService).(service.UserServiceInterface)

	item, err := userService.GetItem(2)

	if err != nil {
		t.Error(err)
	}

	testItem := item.(*entity.User)
	if testItem.ID != 2 {
		t.Error("Id is not equals to 2")
	}

	if testItem.Name != "test user" {
		t.Error("'test user' expected")
	}
}
