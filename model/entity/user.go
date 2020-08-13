package entity

import (
	"github.com/jinzhu/gorm"
	"github.com/zubroide/gorm-crud"
)

type UserFields struct {
	Name string  `binding:"required"`
}

type User struct {
	gorm_crud.InterfaceEntity
	gorm.Model
	*UserFields  `binding:"required"`
}
