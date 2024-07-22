package dto

import "bot/internal/domain/entity"

type Response struct {
	Command string
	Data    interface{}
}

func ErrorMapper(err error) string {
	if err == nil {
		return ""
	}

	switch err.Error() {

	case entity.BadRequset:
		return "Некорректные данные"

	default:
		return err.Error()
	}
}