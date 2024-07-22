package dto

import (
	"bot/internal/domain/entity"

	"gopkg.in/telebot.v3"
)

func ToUser(user *telebot.User) entity.User {
	return entity.User{
		Id:      int(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.Username,
		Role:      "user",
	}
}
