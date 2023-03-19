package telegram

import (
	"github.com/NicoNex/echotron/v3"
)

func (b *Bot) handleAddDiscipline(update *echotron.Update) stateFn {
	_, err := b.SendMessage("Введи название дисциплины", b.chatID, nil)
	if err != nil {
		return b.handleMessage
	}

	return b.handleAddDisciplineName
}

func (b *Bot) handleAddDisciplineName(update *echotron.Update) stateFn {
	b.eventClient.AddDiscipline(update.Message.Text)

	_, err := b.SendMessage("Добавлено", b.chatID, &echotron.MessageOptions{ReplyMarkup: defaultKeyboard(b.user.IsAdmin)})
	if err != nil {
		return b.handleMessage
	}

	return b.handleMessage
}
