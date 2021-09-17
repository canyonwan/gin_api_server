package dto

import "gin_api_server/model"

type UserDto struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

func ConvertUserDto(user model.User) UserDto {
	return UserDto{
		Username: user.Username,
		Phone:    user.Phone,
	}
}
