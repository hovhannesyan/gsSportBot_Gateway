package telegram

import (
	"github.com/NicoNex/echotron/v3"
	"strconv"
)

func (b *Bot) handleSettings(update *echotron.Update) stateFn {
	disciplines := b.connectionClient.GetSet("user", "disciplines", b.user.Id)
	disciplinesMap := make(map[string]bool)

	for _, discipline := range disciplines {
		disciplinesMap[discipline] = true
	}

	if update.CallbackQuery != nil {
		if update.CallbackQuery.Data == "done" {
			_, err := b.DeleteMessage(b.chatID, update.CallbackQuery.Message.ID)
			if err != nil {
				return b.handleMessage
			}
			return b.handleMessage
		}
		disId, _ := strconv.ParseInt(update.CallbackQuery.Data, 10, 64)
		if _, ok := disciplinesMap[update.CallbackQuery.Data]; !ok {
			b.connectionClient.AddToSet("user", "disciplines", b.user.Id, update.CallbackQuery.Data)
			b.connectionClient.AddToSet("discipline", "users", disId, strconv.FormatInt(b.user.Id, 10))
			disciplinesMap[update.CallbackQuery.Data] = true
		} else {
			b.connectionClient.RemoveFromSet("user", "disciplines", b.user.Id, update.CallbackQuery.Data)
			b.connectionClient.RemoveFromSet("discipline", "users", disId, strconv.FormatInt(b.user.Id, 10))
			delete(disciplinesMap, update.CallbackQuery.Data)
		}
		_, err := b.EditMessageReplyMarkup(echotron.NewMessageID(b.chatID, update.CallbackQuery.Message.ID), &echotron.MessageReplyMarkup{
			ReplyMarkup: settingsInlineKeyboard(disciplinesMap, b.eventClient.GetDisciplines()),
		})
		if err != nil {
			return b.handleMessage
		}
		return b.handleSettings
	}

	_, err := b.SendMessage("Выбери интересующие тебя дисциплины", b.chatID, &echotron.MessageOptions{
		ReplyMarkup: settingsInlineKeyboard(disciplinesMap, b.eventClient.GetDisciplines()),
	})
	if err != nil {
		return b.handleMessage
	}

	return b.handleSettings
}
