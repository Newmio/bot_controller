package mongo

import (
	"bot/internal/domain/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type IMongo interface {
	CreateUser(user entity.User) error
	GetServers(limit, offset int64) ([]entity.BotServer, error)
	CreateServer(server entity.BotServer) error
	CreateBot(userId int, login, pass string)error
}

type mongoRepo struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) IMongo {
	return &mongoRepo{db: db}
}
