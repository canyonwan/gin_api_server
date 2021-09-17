package common

import (
	"fmt"
	"gin_api_server/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// InitDB connect mysql
func InitDB() *gorm.DB {
	driverName := "mysql"
	dbUser := "root"
	dbPassword := "1234"
	host := "localhost"
	port := "3306"
	database := "gin_api_db"
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, host, port, database)
	db, openDbError := gorm.Open(driverName, dns)
	if openDbError != nil {
		panic(openDbError.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}