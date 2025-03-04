package logger

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	sentrylogrus "github.com/getsentry/sentry-go/logrus"
	"github.com/sirupsen/logrus"
	"os"
	"runtime/debug"
	"time"
)

type LoggerInterface interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Print(args ...interface{})
}

func NewLogger(sentryHook *sentrylogrus.Hook, level logrus.Level, sentryTimeout int) LoggerInterface {
	logger := NewLoggerWithoutSentry(level)

	if sentryHook != nil {
		logger.AddHook(sentryHook)
		// Flushes before calling os.Exit(1) when using logger.Fatal
		// (else all defers are not called, and Sentry does not have time to send the event)
		logrus.RegisterExitHandler(func() {
			sentryHook.Flush(5 * time.Second)
		})
	}

	return logger
}

func NewSentryHook(dsn string) (*sentrylogrus.Hook, error) {
	sentryLevels := []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
	if dsn != "" {
		sentry.Init(sentry.ClientOptions{
			Dsn:              dsn,
			AttachStacktrace: true,
		})
	}
	sentryHook, err := sentrylogrus.New(sentryLevels, sentry.ClientOptions{
		Dsn:              dsn,
		EnableTracing:    true,
		AttachStacktrace: true,
		TracesSampleRate: 1.0,
		//Debug:            true,
	})
	if err != nil {
		panic(err)
	}
	defer sentryHook.Flush(5 * time.Second)

	return sentryHook, nil
}

func NewLoggerWithoutSentry(level logrus.Level) *logrus.Logger {
	logger := &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		},
		Hooks: make(logrus.LevelHooks),
		Level: level,
	}

	return logger
}

func RecoverPanic() {
	err := recover()

	if err != nil {
		fmt.Println(err)
		fmt.Println(string(debug.Stack()))
		sentry.CurrentHub().Recover(err)
		sentry.Flush(time.Second * 5)
	}
}
