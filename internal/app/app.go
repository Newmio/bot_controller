package app

import (
	m "bot/internal/configs/mongo"
	r "bot/internal/configs/redis"
	"bot/internal/domain/usecase"
	"bot/internal/repository/mongo"
	"bot/internal/repository/redis"
	"bot/internal/transport/telegram"
	"os"
)

func Init() {
	m, err := m.OpenDb()
	if err != nil {
		panic(err)
	}

	r, err := r.OpenDb()
	if err != nil {
		panic(err)
	}

	token, err := os.ReadFile("internal/configs/token.txt")
	if err != nil {
		panic(err)
	}

	mongoRepo := mongo.NewMongo(m.Database("telegram"))
	redisRepo := redis.NewRedis(r)
	usecase := usecase.NewUsecase(mongoRepo, redisRepo)
	handlerBot := telegram.NewHandler(usecase)
	handlerBot.StartBot(string(token))
}
