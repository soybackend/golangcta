package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "cta"
const DB_PASSWORD = "49H4EFYqS5YhKPur"
const DB_NAME = "sellers"
const DB_HOST = "sellers.cfxzweincbgz.us-east-1.rds.amazonaws.com"
const DB_PORT = "3306"

var Db *gorm.DB
func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() (*gorm.DB) {
	var err error
	dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to db: error=%v\n", err)
		return nil
	}

	return db
}