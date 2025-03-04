package main

import (
	"database/sql"
	"github.com/zubroide/go-api-boilerplate/dic"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `binding:"required"`
}

// Up is executed when this migration is applied
func Up_20190516163301(tx *sql.Tx) {
	dic.DB.Migrator().CreateTable(&User{})
}

// Down is executed when this migration is rolled back
func Down_20190516163301(tx *sql.Tx) {
	dic.DB.Migrator().DropTable(&User{})
}
