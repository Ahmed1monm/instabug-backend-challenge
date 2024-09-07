package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbName string) *gorm.DB {
	// dsn := "root:password@tcp(db:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: os.Getenv("DATABASE_URL"),
	}), &gorm.Config{})
	if err != nil {
		panic(">>>> failed to connect database")
	}
	log.Println(">>>> connected to database")
	DB = db
	return db
}

func Migrate(dst ...interface{}) {
	DB.AutoMigrate(dst...)
}
