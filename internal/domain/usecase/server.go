package usecase

import "bot/internal/domain/entity"

func (s *usecase) CreateServer(server entity.BotServer) error{
	return s.r.CreateServer(server)
}

func (s *usecase) GetServers(limit, offset int64) ([]entity.BotServer, error) {
	return s.r.GetServers(limit, offset)
}
