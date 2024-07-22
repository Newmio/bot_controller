package telegram

import (
	"bot/internal/domain/usecase"
	"bot/internal/dto"
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
			Description: "Список ботов",
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
	resp, err := h.s.OnText(dto.ToUser(c.Sender()), c.Text())
	if err != nil {
		return c.Send(dto.ErrorMapper(err))
	}

	switch resp.Command {

	case "/addbot":
		return h.addBotResponse(c)
	}

	return c.Send("Ошибка обработки комманды (*_*)\nПовторите попытку либо обратитесь в поддержку если ошибка повториться")
}
