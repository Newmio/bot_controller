package telegram

import (
	"bot/pkg"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/telebot.v3"
)

func (h *handler) botsCommand(c telebot.Context) error {

	inlineMenu, err := h.generateBotsButtons(0)
	if err != nil {
		return c.Send(err)
	}

	return c.Send("Ваши боты:", inlineMenu)
}

func (h *handler) botsButton(c telebot.Context) error {
	offset, err := strconv.Atoi(strings.Split(c.Data(), "|")[1])
	if err != nil {
		return c.Send(err.Error())
	}

	inlineMenu, err := h.generateBotsButtons(offset)
	if err != nil {
		return c.Send(err)
	}

	return c.Send("Ваши боты:", inlineMenu)
}

func (h *handler) generateBotsButtons(offset int) (*telebot.ReplyMarkup, error) {
	if offset < 0 {
		offset = 0
	}

	bots, err := h.s.GetBots(9, int64(offset))
	if err != nil {
		return nil, pkg.Trace(err)
	}

	var rows []telebot.Row
	var buttons []telebot.Btn

	for i, bot := range bots {
		buttons = append(buttons, telebot.Btn{
			Text:   bot.Login,
			Unique: bot.Login,
		})

		if (i+1)%3 == 0 {
			rows = append(rows, buttons)
			buttons = nil
		}
	}

	if len(buttons) > 0 {
		rows = append(rows, buttons)
	}

	rows = append(rows, []telebot.Btn{
		{
			Text:   "←",
			Unique: "bots_previous",
			Data:   fmt.Sprint(offset - 9),
		},
		{
			Text:   "→",
			Unique: "bots_next",
			Data:   fmt.Sprint(offset + 9),
		},
	})

	inlineMenu := &telebot.ReplyMarkup{}
	inlineMenu.Inline(rows...)

	return inlineMenu, nil
}

func (h *handler) addBot(c telebot.Context) error {
	err := c.Send("Вы перешли в опцию создания бота")
	if err != nil {
		return c.Send(err)
	}

	return c.Send(h.s.Addbot(int(c.Sender().ID), "").Error())
}

func (h *handler) addBotResponse(c telebot.Context) error {
	return c.Send("Бот успешно добавлен!\nЕго можно настроить через:\n1. Найти бота в списке ботов - /bots\n2. Сразу зайти в бота - /bots steamlogin")
}
