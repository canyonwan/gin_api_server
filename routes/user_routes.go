package routes

import (
	"gin_api_server/controller"
	"gin_api_server/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) *gin.Engine {
	r.POST("api/user/register", controller.Register)
	r.POST("api/user/login", controller.Login)
	r.GET("api/user/profile", middleware.AuthMiddleware(), controller.GetUserProfile)
	return r
}
