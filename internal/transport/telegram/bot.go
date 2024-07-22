package telegram

import (
	"gopkg.in/telebot.v3"
)

func (h *handler) botList(c telebot.Context) error {
	return nil
}

func (h *handler) addBot(c telebot.Context) error {
	return c.Send("Введите через заптую:\n1. Стим логин\n2. Стим пароль\nПример: steamlogin, steampassword")
}

func (h *handler) addBotResponse(c telebot.Context) error {
	return c.Send("Бот успешно добавлен!\nЕго можно настроить через:\n1. Найти бота в списке ботов - /bots\n2. Сразу зайти в бота - /bots steamlogin")
}
