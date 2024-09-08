package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbName string) *gorm.DB {
	// dsn := "root:password@tcp(db:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	time.Sleep(60 * time.Second)

	dsn := "root:password@tcp(db:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	var db *gorm.DB
	var err error

	for db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{}); err != nil; {
		fmt.Println("Error connecting to database: ", err)
		time.Sleep(5 * time.Second)
		fmt.Println("Retrying...")
	}

	log.Println(">>>> connected to database")
	DB = db
	return db
}

func Migrate(dst ...interface{}) {
	DB.AutoMigrate(dst...)
}
