package utils

import (
	"gin_api_server/model"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)

// RandomName 随机名
func RandomName(n int) string {
	letters := []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLMNBVCXZ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// IsPhoneExist 判断手机号是否存在
func IsPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
