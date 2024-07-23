package telegram

import (
	"gopkg.in/telebot.v3"
)

func (h *handler) botList(c telebot.Context) error {
	return nil
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
