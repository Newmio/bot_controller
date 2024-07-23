package usecase

import (
	"bot/internal/domain/entity"
	"bot/internal/dto"
	"bot/pkg"
	"fmt"
)

func (s *usecase) OnText(user entity.User, text string) (dto.Response, error) {
	var resp dto.Response

	command, err := s.r.GetSessionCommand(user.Id)
	if err != nil {
		return resp, pkg.Trace(err)
	}

	switch command {

	case entity.CommandAddBot:
		resp.Command = entity.CommandAddBot

		if err := s.Addbot(user.Id, text); err != nil {
			return resp, err
		}

	default:
		return resp, fmt.Errorf("%s, не понимаю комманды", user.UserName)
	}

	return resp, nil
}

func (s *usecase) SetSessionCommand(userId int, command string) error{
	return s.r.SetSessionCommand(userId, command)
}