package controller

import (
	"gin_api_server/common"
	"gin_api_server/model"
	"gin_api_server/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// Register 用户注册
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	username := ctx.PostForm("username")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(username) == 0 {
		username = utils.RandomName(10)
	}
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "手机号必须11位!",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "密码必须大于6位",
		})
		return
	}
	// 判断手机号是否存在
	if isPhoneExist(DB, phone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "手机号已存在, 请更换其他手机号",
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "加密失败",
		})
		return
	}
	// 创建用户
	newUser := model.User{
		Username: username,
		Phone:    phone,
		Password: string(hashedPassword),
	}
	DB.Create(&newUser)
	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "注册成功",
	})
}

// Login 用户登录
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "手机号必须为11位",
		})
		return
	}
	// 判断手机是否存在
	var user model.User
	DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "用户不存在, 滚去注册好嘛~",
		})
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "密码必须大于6位",
		})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "密码错误",
		})
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"message": "系统内部错误",
		})
		log.Printf("token generate error: %v", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}
func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
