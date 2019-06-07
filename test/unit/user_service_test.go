package unit

import (
	"go-api-boilerplate/dic"
	"go-api-boilerplate/model/entity"
	"go-api-boilerplate/model/repository"
	"go-api-boilerplate/model/service"
	"github.com/mnvx/di"
	"testing"
)

type UserRepositoryMock struct {
	repository.CrudRepositoryInterface
	model entity.InterfaceEntity // Dynamic typing
}

func NewUserRepositoryMock() repository.UserRepositoryInterface {
	var model entity.User
	repo := &UserRepositoryMock{model: model}
	return repo
}

func (c UserRepositoryMock) Find(id uint) (entity.InterfaceEntity, error) {
	item := &entity.User{UserFields: &entity.UserFields{Name: "test user"}}
	item.ID = 2
	return item, nil
}

func TestGetItem(t *testing.T) {
	// It is not good way, because we cannot override provided type with another.
	// So we must specify here all services
	dic.InitBuilder()
	dic.Builder.Set(di.Def{
		Name: dic.UserRepository,
		Build: func(ctn di.Container) (interface{}, error) {
			return NewUserRepositoryMock(), nil
		},
	})
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
