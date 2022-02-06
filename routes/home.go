package routes

import (
	"gin_api_server/controller"
	"github.com/gin-gonic/gin"
)

func HomeRouters(g *gin.Engine) *gin.Engine {
	g.GET("/api/home/swiperList", controller.GetHomeSwiperList)
	g.GET("/api/home/goodsList", controller.GetHomeGoods)
	return g
}
