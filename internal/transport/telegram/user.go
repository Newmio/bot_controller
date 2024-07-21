package telegram

import (
	"bot/internal/dto"
	"bot/pkg"

	"gopkg.in/telebot.v3"
)

func (h *handler) start(c telebot.Context) error {
	if err := h.s.CreateUser(dto.ToUser(c.Sender())); err != nil {
		return pkg.Trace(err)
	}

	return c.Send("Hello!")
}
