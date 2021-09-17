package middleware

import (
	"gin_api_server/common"
	"gin_api_server/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取Header
		tokenString := ctx.GetHeader("Authorization")

		// 验证前端传入的token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			// 权限不足,将当前请求抛弃
			ctx.Abort()
			return
		}

		// 验证通过,提取token
		tokenString = tokenString[7:]

		// 解析token
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			ctx.Abort()
			return
		}

		// 验证token通过后, 获取claims中的用户id
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 判断用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			ctx.Abort()
			return
		}
		// 用户存在, 将用户信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
