package database

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB { // OOP constructor
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	// github.com/denisenkom/go-mssqldb
	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error...")
		return nil
	}
	return db
}
