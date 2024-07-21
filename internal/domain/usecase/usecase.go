package usecase

import (
	"bot/internal/domain/entity"
	"bot/internal/repository/mongo"
)

type IUsecase interface {
	CreateUser(user entity.User) error
	GetServers(limit, offset int64) ([]entity.BotServer, error)
	CreateServer(server entity.BotServer) error
}

type usecase struct {
	r mongo.IMongo
}

func NewUsecase(r mongo.IMongo) IUsecase {
	return &usecase{r: r}
}
