package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDb() *gorm.DB {
	dsn := viper.GetString("DB_URL")

	logMode := logger.Silent
	if viper.GetBool("DB_LOG_MODE") {
		logMode = logger.Info
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database instance")
	}

	sqlDB.SetMaxIdleConns(viper.GetInt("DB_MAX_CONNECTIONS"))
	sqlDB.SetMaxOpenConns(viper.GetInt("DB_MAX_CONNECTIONS"))

	return db
}
