package app

import (
	m "bot/internal/configs/mongo"
	"bot/internal/domain/usecase"
	"bot/internal/repository/mongo"
	"bot/internal/transport/telegram"
	"fmt"
	"os"
)

func Init() {
	db, err := m.OpenDb()
	if err != nil {
		panic(err)
	}

	token, err := os.ReadFile("internal/configs/token.txt")
	if err != nil {
		panic(err)
	}

	mongoRepo := mongo.NewMongo(db.Database("telegram"))
	usecase := usecase.NewUsecase(mongoRepo)
	handlerBot := telegram.NewHandler(usecase)
	fmt.Println("start")
	handlerBot.StartBot(string(token))
}
