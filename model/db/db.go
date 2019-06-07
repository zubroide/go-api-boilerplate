package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func NewDb() *gorm.DB {
	dbUrl := viper.GetString("DB_URL")
	db, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	return db
}
