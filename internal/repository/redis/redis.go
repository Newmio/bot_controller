package redis

import (
	"github.com/redis/go-redis/v9"
)

type IRedis interface {
	GetSessionCommand(userId int) (string, error)
	SetSessionCommand(userId int, command string) error
	SetCommandActions(userId int, command string, actions map[string]string) error
	GetCommandActions(userId int, command string) (map[string]string, error)
	DeleteCommandActions(userId int, command string) error
}

type redisRepo struct {
	db *redis.ClusterClient
}

func NewRedis(db *redis.ClusterClient) IRedis {
	return &redisRepo{db: db}
}
