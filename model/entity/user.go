package entity

import (
	"github.com/jinzhu/gorm"
)

type InterfaceEntity interface {
}

type UserFields struct {
	Name string  `binding:"required"`
}

type User struct {
	InterfaceEntity
	gorm.Model
	*UserFields  `binding:"required"`
}
