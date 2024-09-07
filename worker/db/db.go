package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbName string) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: os.Getenv("DATABASE_URL"),
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	return db
}

func Migrate(dst ...interface{}) {
	DB.AutoMigrate(dst...)
}
