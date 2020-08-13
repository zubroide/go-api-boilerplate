package main

import (
	"github.com/zubroide/go-api-boilerplate/command"
	"github.com/zubroide/go-api-boilerplate/dic"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/spf13/viper"
	"os"
)

func readConfig() {
	var err error

	viper.SetConfigFile("base.env")
	viper.SetConfigType("props")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("WARNING: file .env not found")
	} else {
		viper.SetConfigFile(".env")
		viper.SetConfigType("props")
		err = viper.MergeInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Override config parameters from environment variables if specified
	for _, key := range viper.AllKeys() {
		viper.BindEnv(key)
	}
}

func main() {
	readConfig()

	dic.InitContainer()

	client := dic.Container.Get(dic.RavenClient).(*raven.Client)
	if client != nil {
		client.CapturePanicAndWait(func() {
			command.Execute()
		}, nil)
	} else {
		command.Execute()
	}
}
