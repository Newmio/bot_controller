package telegram

import (
	"bot/internal/domain/usecase"
	"time"

	"gopkg.in/telebot.v3"
)

type handler struct {
	s usecase.IUsecase
}

func NewHandler(s usecase.IUsecase) *handler {
	return &handler{s: s}
}

func (h *handler) StartBot(token string) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 2 * time.Second},
	})
	if err != nil {
		panic(err)
	}

	err = bot.SetCommands([]telebot.Command{
		{
			Text:        "/start",
			Description: "Старт",
		},
		{
			Text:        "/bots",
			Description: "Список всех ботов",
		},
		{
			Text:        "/addbot",
			Description: "Добавить бота",
		},
		{
			Text:        "/servers",
			Description: "Список серверов",
		},
	})
	if err != nil {
		panic(err)
	}

	bot.Handle("/start", h.start)
	bot.Handle("/bots", h.botList)
	bot.Handle("/addbot", h.addBot)
	bot.Handle("/servers", h.servers)
	bot.Handle(telebot.OnText, h.onText)

	bot.Start()
}

func (h *handler) onText(c telebot.Context) error {
	return nil
}
