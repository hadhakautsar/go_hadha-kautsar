package config

import (
	"fmt"
	"section-20/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DB_USER = "root"
const DB_PASS = ""
const DB_HOST = "localhost"
const DB_PORT = "3306"
const DB_NAME = "gorm"

const DB_USER_TEST = "root"
const DB_PASS_TEST = ""
const DB_HOST_TEST = "localhost"
const DB_PORT_TEST = "3306"
const DB_NAME_TEST = "gorm"

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&model.Users{})
}

func InitDBTest() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER_TEST, DB_PASS_TEST, DB_HOST_TEST, DB_PORT_TEST, DB_NAME_TEST)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&model.Users{})
	DB.AutoMigrate(&model.Users{})
}
