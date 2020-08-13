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

	db.DB().SetMaxIdleConns(viper.GetInt("DB_MAX_CONNECTIONS"))
	db.DB().SetMaxOpenConns(viper.GetInt("DB_MAX_CONNECTIONS"))
	db.LogMode(viper.GetBool("DB_LOG_MODE"))

	return db
}
