package main

import (
	"gin_api_server/common"
	"gin_api_server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	// 日志颜色化
	gin.ForceConsoleColor()

	// 连接数据库
	db := common.InitDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// 初始化gin引擎
	r := gin.Default()

	// 注册路由
	//r = routes.UserRoutes(r)
	r = routes.HomeRouters(r)

	// 启动gin
	panic(r.Run())
}
