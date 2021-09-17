package main

import (
	"fmt"
	"gin_api_server/common"
	"gin_api_server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	gin.ForceConsoleColor()

	db := common.InitDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	r := gin.Default()
	r = routes.UserRoutes(r)

	if err := r.Run(); err != nil {
		fmt.Println("failed gin start:", err)
	}
}
