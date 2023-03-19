package telegram

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/hovhannesyan/gsSportBot_Gateway/pkg/models"
)

func (b *Bot) handleStart(update *echotron.Update) stateFn {
	if b.userClient.GetUsersById(update.Message.From.ID) != nil {
		_, err := b.SendMessage("Ты уже зареган", b.chatID, nil)
		if err != nil {
			return b.handleMessage
		}
		return b.handleMessage
	}

	_, err := b.SendMessage("Привет!\nКак тебя зовут? (Имя Фамилия)", b.chatID, nil)
	if err != nil {
		return b.handleMessage
	}

	b.user = &models.User{Id: update.Message.From.ID, Username: update.Message.From.Username}

	return b.handleName
}

func (b *Bot) handleName(update *echotron.Update) stateFn {
	b.user.Name = update.Message.Text

	_, err := b.SendMessage("Приятно познакомится!\n\nА теперь введи свой номер телефона в формате\n+420 XXX XXX XXX", b.chatID, nil)
	if err != nil {
		return b.handleMessage
	}

	return b.handlePhone
}

func (b *Bot) handlePhone(update *echotron.Update) stateFn {
	b.user.Phone = update.Message.Text

	_, err := b.SendMessage("Если ты админ, то отправь мне пин", b.chatID, nil)
	if err != nil {
		return b.handleMessage
	}

	return b.handleAdminProp
}

func (b *Bot) handleAdminProp(update *echotron.Update) stateFn {
	if update.Message.Text == "allez" {
		b.user.IsAdmin = true
	} else {
		b.user.IsAdmin = false
	}

	b.userClient.AddUser(b.user.Id, b.user.Username, b.user.Name, b.user.Phone, b.user.IsAdmin)

	_, err := b.SendMessage("Отлично\nДля взаимодействия со мной у тебя есть TODO кнопки:\n1. Настройки: нажав на нее ты можешь настроить свои уведомления (TODO)",
		b.chatID,
		&echotron.MessageOptions{ReplyMarkup: defaultKeyboard(b.user.IsAdmin)})
	if err != nil {
		return b.handleMessage
	}
	return b.handleMessage
}
