package logger

import (
	"fmt"
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

type LoggerInterface interface {
	Error(args ...interface{})
}

func NewLogger(client *raven.Client) LoggerInterface {
	logger := logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{ForceColors: true},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}

	if client != nil {
		hook, err := logrus_sentry.NewWithClientSentryHook(client, []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		})
		timeout := viper.GetInt("SENTRY_TIMEOUT")
		hook.Timeout = time.Duration(timeout) * time.Second
		hook.StacktraceConfiguration.Enable = true

		if err == nil {
			logger.Hooks.Add(hook)
		}
	}

	return &logger
}

func NewRavenClient() *raven.Client {
	dsn := viper.Get("SENTRY_DSN")
	if dsn == nil {
		return nil
	}

	client, err := raven.New(dsn.(string))
	if err != nil {
		fmt.Println("Fatal")
		fmt.Println(err)
	}
	return client
}
