package dic

import (
	"go-api-boilerplate/controller"
	"go-api-boilerplate/logger"
	"go-api-boilerplate/model/db"
	"go-api-boilerplate/model/repository"
	"go-api-boilerplate/model/service"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/mnvx/di"
)

var Builder *di.Builder
var Container di.Container

const RavenClient = "raven_client"
const Logger = "logger"
const Db = "db"
const UserRepository = "repository.user"
const UserService = "service.user"
const UserController = "controller.user"

func InitContainer() di.Container {
	builder := InitBuilder()
	Container = builder.Build()
	return Container
}

func InitBuilder() *di.Builder {
	Builder, _ = di.NewBuilder()
	RegisterServices(Builder)
	return Builder
}

func RegisterServices(builder *di.Builder) {
	builder.Add(di.Def{
		Name: RavenClient,
		Build: func(ctn di.Container) (interface{}, error) {
			return logger.NewRavenClient(), nil
		},
	})

	builder.Add(di.Def{
		Name: Logger,
		Build: func(ctn di.Container) (interface{}, error) {
			return logger.NewLogger(ctn.Get(RavenClient).(*raven.Client)), nil
		},
	})

	builder.Add(di.Def{
		Name: Db,
		Build: func(ctn di.Container) (interface{}, error) {
			return db.NewDb(), nil
		},
	})

	builder.Add(di.Def{
		Name: UserRepository,
		Build: func(ctn di.Container) (interface{}, error) {
			return repository.NewUserRepository(ctn.Get(Db).(*gorm.DB), ctn.Get(Logger).(*logger.Logger)), nil
		},
	})

	builder.Add(di.Def{
		Name: UserService,
		Build: func(ctn di.Container) (interface{}, error) {
			return service.NewUserService(ctn.Get(UserRepository).(repository.UserRepositoryInterface), ctn.Get(Logger).(*logger.Logger)), nil
		},
	})

	builder.Add(di.Def{
		Name: UserController,
		Build: func(ctn di.Container) (interface{}, error) {
			return controller.NewUserController(ctn.Get(UserService).(service.UserServiceInterface), ctn.Get(Logger).(*logger.Logger)), nil
		},
	})
}
