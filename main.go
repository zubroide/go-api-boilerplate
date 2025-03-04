package main

import (
	"github.com/zubroide/go-api-boilerplate/command"
	"github.com/zubroide/go-api-boilerplate/dic"
)

func main() {
	dic.ReadConfig()
	dic.InitContainer()
	command.Execute()
}
