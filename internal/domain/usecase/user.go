package usecase

import (
	"bot/internal/domain/entity"
)

func (u *usecase) CreateUser(user entity.User) error {
	return u.r.CreateUser(user)
}
