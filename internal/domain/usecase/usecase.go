package usecase

import (
	"bot/internal/domain/entity"
	"bot/internal/dto"
	"bot/internal/repository/mongo"
	"bot/internal/repository/redis"
	"bot/pkg"
	"fmt"
)

type IUsecase interface {
	CreateUser(user entity.User) error
	GetServers(limit, offset int64) ([]entity.BotServer, error)
	CreateServer(server entity.BotServer) error
	OnText(user entity.User, text string) (dto.Response, error)
	SetSessionCommand(userId int, command string) error
	Addbot(userId int, arg string) error
}

type usecase struct {
	m mongo.IMongo
	r redis.IRedis
}

func NewUsecase(m mongo.IMongo, r redis.IRedis) IUsecase {
	return &usecase{m: m, r: r}
}

func (s *usecase) sessionCommandMapper(userId int, template map[string]string, command, argument string) (map[string]string, error) {
	actions, err := s.r.GetCommandActions(userId, command)
	if err != nil {
		return nil, pkg.Trace(err)
	}

	if len(actions) == 0 {
		for key := range template {
			actions[key] = ""
		}
	}

	for key, value := range actions {
		if value == "wait" && argument != "" {
			actions[key] = argument
			break
		}
	}

	for key, value := range template {
		if v, ok := actions[key]; !ok || v == "" {

			actions[key] = "wait"

			if err := s.r.SetCommandActions(userId, command, actions); err != nil {
				return nil, pkg.Trace(err)
			}

			return nil, fmt.Errorf(value)
		}
	}

	return actions, nil
}
