package dic

import (
	sentrylogrus "github.com/getsentry/sentry-go/logrus"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zubroide/go-api-boilerplate/controller"
	"github.com/zubroide/go-api-boilerplate/logger"
	"github.com/zubroide/go-api-boilerplate/model/db"
	"github.com/zubroide/go-api-boilerplate/model/service"
	"gorm.io/gorm"
	"strings"
)

var Builder *di.Builder
var Container di.Container

const SentryHook = "SentryHook"
const Logger = "logger"
const Db = "db"
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
		Name: SentryHook,
		Build: func(ctn di.Container) (interface{}, error) {
			dsn := viper.GetString("SENTRY_DSN")
			if dsn == "" {
				var sh *sentrylogrus.Hook
				return sh, nil
			}
			return logger.NewSentryHook(dsn)
		},
	})

	builder.Add(di.Def{
		Name: Logger,
		Build: func(ctn di.Container) (interface{}, error) {
			level, _ := logrus.ParseLevel(strings.ToLower(
				viper.GetString("LOG_LEVEL"),
			))
			return logger.NewLogger(
				ctn.Get(SentryHook).(*sentrylogrus.Hook),
				//nil,
				level,
				viper.GetInt("SENTRY_TIMEOUT"),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: Db,
		Build: func(ctn di.Container) (interface{}, error) {
			return db.NewDb(), nil
		},
	})

	builder.Add(di.Def{
		Name: UserService,
		Build: func(ctn di.Container) (interface{}, error) {
			return service.NewUserService(ctn.Get(Db).(*gorm.DB), ctn.Get(Logger).(logger.LoggerInterface)), nil
		},
	})

	builder.Add(di.Def{
		Name: UserController,
		Build: func(ctn di.Container) (interface{}, error) {
			return controller.NewUserController(ctn.Get(UserService).(service.UserServiceInterface), ctn.Get(Logger).(logger.LoggerInterface)), nil
		},
	})
}
