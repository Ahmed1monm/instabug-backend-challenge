package db

import (
	"fmt"
	"os"

	// "time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "worker/models"
)

var DB *gorm.DB

func InitDB(dbName string) *gorm.DB {
	// time.Sleep(60 * time.Second)

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	var db *gorm.DB
	var err error

	for db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil; {
		fmt.Println("Error connecting to database: ", err)
		panic("Error connecting to database")
	}

	fmt.Println("Connected to database")

	DB = db
	return db
}

func Migrate(dst ...interface{}) {
	DB.AutoMigrate(dst...)
}
