
package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
}

// Up is executed when this migration is applied
func Up_20190516163301(txn *gorm.DB) {
	txn.CreateTable(&User{})
}

// Down is executed when this migration is rolled back
func Down_20190516163301(txn *gorm.DB) {
	txn.DropTable(&User{})
}
