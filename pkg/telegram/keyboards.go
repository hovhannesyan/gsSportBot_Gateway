package telegram

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/hovhannesyan/gsSportBot_Gateway/pkg/event/pb"
	"strconv"
)

func defaultKeyboard(isAdmin bool) echotron.ReplyKeyboardMarkup {
	var keyboard echotron.ReplyKeyboardMarkup

	if isAdmin {
		keyboard = echotron.ReplyKeyboardMarkup{
			Keyboard: [][]echotron.KeyboardButton{{
				{Text: "Тренировки"}, {Text: "Настройки"},
			}, {
				{Text: "Добавить тренировку"}, {Text: "Добавить дисциплину"},
			}},
		}
	} else {
		keyboard = echotron.ReplyKeyboardMarkup{
			Keyboard: [][]echotron.KeyboardButton{{
				{Text: "Тренировки"}, {Text: "Настройки"},
			}},
		}
	}

	return keyboard
}

func settingsInlineKeyboard(info map[string]bool, data []*pb.DisciplineData) echotron.InlineKeyboardMarkup {
	var keyboard echotron.InlineKeyboardMarkup
	if len(data)%2 == 0 {
		for i := 0; i < len(data); i += 2 {
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []echotron.InlineKeyboardButton{
				{Text: exists(info, data[i].Id) + data[i].Name, CallbackData: strconv.FormatInt(data[i].Id, 10)},
				{Text: exists(info, data[i+1].Id) + data[i+1].Name, CallbackData: strconv.FormatInt(data[i+1].Id, 10)},
			})
		}
	} else {
		for i := 0; i < len(data)-1; i += 2 {
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []echotron.InlineKeyboardButton{
				{Text: exists(info, data[i].Id) + data[i].Name, CallbackData: strconv.FormatInt(data[i].Id, 10)},
				{Text: exists(info, data[i+1].Id) + data[i+1].Name, CallbackData: strconv.FormatInt(data[i+1].Id, 10)},
			})
		}
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []echotron.InlineKeyboardButton{
			{Text: exists(info, data[len(data)-1].Id) + data[len(data)-1].Name, CallbackData: strconv.FormatInt(data[len(data)-1].Id, 10)},
		})
	}
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []echotron.InlineKeyboardButton{
		{Text: "Готово", CallbackData: "done"},
	})
	return keyboard
}

func exists(info map[string]bool, id int64) string {
	if _, ok := info[strconv.FormatInt(id, 10)]; ok {
		return "■ "
	}
	return "□ "
}

func chooseDisciplineInlineKeyboard(data []*pb.DisciplineData) echotron.InlineKeyboardMarkup {
	var keyboard echotron.InlineKeyboardMarkup
	if len(data)%2 == 0 {
		for i := 0; i < len(data); i += 2 {
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []echotron.InlineKeyboardButton{
				{Text: data[i].Name, CallbackData: strconv.FormatInt(data[i].Id, 10)},
				{Text: data[i+1].Name, CallbackData: strconv.FormatInt(data[i+1].Id, 10)},
			})
		}
	} else {
		for i := 0; i < len(data)-1; i += 2 {
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []echotron.InlineKeyboardButton{
				{Text: data[i].Name, CallbackData: strconv.FormatInt(data[i].Id, 10)},
				{Text: data[i+1].Name, CallbackData: strconv.FormatInt(data[i+1].Id, 10)},
			})
		}
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []echotron.InlineKeyboardButton{
			{Text: data[len(data)-1].Name, CallbackData: strconv.FormatInt(data[len(data)-1].Id, 10)},
		})
	}
	return keyboard
}

func eventInlineKeyboard(eventId int64, signedUp bool) echotron.InlineKeyboardMarkup {
	if signedUp {
		return echotron.InlineKeyboardMarkup{
			InlineKeyboard: [][]echotron.InlineKeyboardButton{{
				{Text: "Отписаться", CallbackData: strconv.FormatInt(eventId, 10)},
			}},
		}
	}
	return echotron.InlineKeyboardMarkup{
		InlineKeyboard: [][]echotron.InlineKeyboardButton{{
			{Text: "Записаться", CallbackData: strconv.FormatInt(eventId, 10)},
		}},
	}
}
