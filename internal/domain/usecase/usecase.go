package usecase

import (
	"bot/internal/domain/entity"
	"bot/internal/dto"
	"bot/internal/repository/mongo"
	"bot/internal/repository/redis"
)

type IUsecase interface {
	CreateUser(user entity.User) error
	GetServers(limit, offset int64) ([]entity.BotServer, error)
	CreateServer(server entity.BotServer) error
	OnText(user entity.User, text string) (dto.Response, error)
}

type usecase struct {
	m mongo.IMongo
	r redis.IRedis
}

func NewUsecase(m mongo.IMongo, r redis.IRedis) IUsecase {
	return &usecase{m: m, r: r}
}
