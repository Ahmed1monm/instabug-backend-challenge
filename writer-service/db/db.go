package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"writer-service/models"
)

var DB *gorm.DB

func InitDB(dbName string) *gorm.DB {
	dsn := "root:password@tcp(db:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	return db
}

func Migrate() {
	DB.AutoMigrate(&models.Application{}, &models.Chat{}, &models.Message{})
}
