package telegram

import (
	"bot/pkg"

	"gopkg.in/telebot.v3"
)

func (h *handler) botList(c telebot.Context) error {
	return nil
}

func (h *handler) addBot(c telebot.Context) error {
	if err := c.Send("Введите через заптую:\n1. Логин стим\n2. Пароль стим\nНапример: steamlogin, steampassword"); err != nil {
		return pkg.Trace(err)
	}
	return nil
}
