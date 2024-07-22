package usecase

import (
	"bot/internal/domain/entity"
	"bot/internal/dto"
	"bot/pkg"
)

func (s *usecase) OnText(user entity.User, text string) (dto.Response, error) {
	var resp dto.Response

	session, err := s.r.GetSession(user.Id)
	if err != nil {
		return resp, pkg.Trace(err)
	}

	switch session.Command {

	case "/addbot":
		resp.Command = "addbot"

		if err := s.addbot(user.Id, text); err != nil {
			return resp, pkg.Trace(err)
		}
	}

	return resp, nil
}
