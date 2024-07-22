package usecase

import (
	"bot/internal/domain/entity"
	"fmt"
	"strings"
)

func (s *usecase) addbot(userId int, args string)error{
	parts := strings.Split(strings.ReplaceAll(args, " ", ""), ",")

	if len(parts) != 2 || len([]rune(parts[1])) < 8{
		return fmt.Errorf(entity.BadRequset)
	}

	return s.m.CreateBot(userId, parts[0], parts[1])
}

func (s *usecase) CreateServer(server entity.BotServer) error{
	return s.m.CreateServer(server)
}

func (s *usecase) GetServers(limit, offset int64) ([]entity.BotServer, error) {
	return s.m.GetServers(limit, offset)
}
