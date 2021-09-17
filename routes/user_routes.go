package routes

import (
	"gin_api_server/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) *gin.Engine  {
	r.POST("api/user/register", controller.Register)
	r.POST("api/user/login", controller.Login)
	return r
}