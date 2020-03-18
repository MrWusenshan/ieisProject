package dto

import "irisProject/model"

//过滤后的消息结构体
type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

//过滤函数
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
