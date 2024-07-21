package telegram

import (
	"bot/pkg"

	"gopkg.in/telebot.v3"
)

func (h *handler) servers(c telebot.Context) error {
	limit, offset := 5, 0

	servers, err := h.s.GetServers(int64(limit), int64(offset))
	if err != nil {
		return pkg.Trace(err)
	}

	inlineMenu := &telebot.ReplyMarkup{}

	var rows []telebot.Row
	for i := range servers{
		var row []telebot.Btn

		for j := i; j < i+3 && j < len(servers); j++ {
			row = append(row, telebot.Btn{
				Text:     servers[j].Name,
				Unique:   servers[j].IP,
			})
		}

		rows = append(rows, row)
	}
	inlineMenu.Inline(rows...)

	return c.Send("Сервера", inlineMenu)
}
