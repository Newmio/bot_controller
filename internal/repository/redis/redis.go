package redis

import (
	"bot/internal/domain/entity"

	"github.com/redis/go-redis/v9"
)

type IRedis interface {
	CreateSession(session entity.Session) error
	GetSession(userId int) (entity.Session, error)
}

type redisRepo struct {
	db *redis.ClusterClient
}

func NewRedis(db *redis.ClusterClient) IRedis {
	return &redisRepo{db: db}
}
