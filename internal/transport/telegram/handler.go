package telegram

import (
	"bot/internal/domain/entity"
	"bot/internal/domain/usecase"
	"bot/internal/dto"
	"bot/pkg"
	"strings"
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
		Poller: &telebot.LongPoller{Timeout: time.Second},
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

	bot.Use(h.setCommandSession)

	bot.Handle("/start", h.start)
	bot.Handle(entity.CommandBots, h.botsCommand)
	bot.Handle(telebot.OnCallback, h.onCallback)
	bot.Handle(entity.CommandAddBot, h.addBot)
	bot.Handle("/servers", h.servers)
	bot.Handle(telebot.OnText, h.onText)

	bot.Start()
}

func (h *handler) onCallback(c telebot.Context) error {
	switch {

	case strings.HasPrefix(strings.TrimSpace(c.Data()), "bots_"):
		return h.botsButton(c)

	default:
		return c.Send("Неизвестная кнопка")
	}
}

func (h *handler) setCommandSession(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {

		if strings.HasPrefix(c.Text(), "/") && !strings.Contains(entity.IgnoreCommandsForSession, c.Text()) {
			if err := h.s.SetSessionCommand(int(c.Sender().ID), c.Text()); err != nil {
				return pkg.Trace(err)
			}
		}

		return next(c)
	}
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
