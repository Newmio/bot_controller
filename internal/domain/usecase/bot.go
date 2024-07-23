package usecase

import (
	"bot/internal/domain/entity"
	"bot/pkg"
)

func (s *usecase) GetBots(limit, offset int64) ([]entity.Bot, error) {
	return s.m.GetBots(limit, offset)
}

func (s *usecase) Addbot(userId int, arg string) error {
	actions, err := s.sessionCommandMapper(userId, entity.AddBotTemplate, entity.CommandAddBot, arg)
	if err != nil {
		return err
	}

	if err := s.m.CreateBot(entity.Bot{UserId: userId, Login: actions["login"], Pass: actions["pass"]}); err != nil {
		return pkg.Trace(err)
	}

	return s.r.DeleteCommandActions(userId, entity.CommandAddBot)
}
