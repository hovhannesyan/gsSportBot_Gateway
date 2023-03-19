package telegram

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/hovhannesyan/gsSportBot_Gateway/pkg/models"
	"strconv"
)

func (b *Bot) handleAddEvent(update *echotron.Update) stateFn {
	_, err := b.SendMessage("Выбери дисциплину", b.chatID, &echotron.MessageOptions{ReplyMarkup: chooseDisciplineInlineKeyboard(b.eventClient.GetDisciplines())})
	if err != nil {
		return b.handleMessage
	}

	b.event = &models.Event{}

	return b.handleAddEventDisciplineId
}

func (b *Bot) handleAddEventDisciplineId(update *echotron.Update) stateFn {
	disId, _ := strconv.ParseInt(update.CallbackQuery.Data, 10, 64)
	b.event.DisciplineId = disId

	_, err := b.DeleteMessage(b.chatID, update.CallbackQuery.Message.ID)
	if err != nil {
		return b.handleMessage
	}

	_, err = b.SendMessage("Введи место тренировки", b.chatID, nil)
	if err != nil {
		return b.handleMessage
	}

	return b.handleAddEventPlace
}

func (b *Bot) handleAddEventPlace(update *echotron.Update) stateFn {
	b.event.Place = update.Message.Text

	_, err := b.SendMessage("Введи дату тренировки", b.chatID, nil)
	if err != nil {
		return b.handleMessage
	}

	return b.handleAddEventDate
}

func (b *Bot) handleAddEventDate(update *echotron.Update) stateFn {
	b.event.Date = update.Message.Text

	_, err := b.SendMessage("Введи время начала тренировки", b.chatID, nil)
	if err != nil {
		return b.handleMessage
	}

	return b.handleAddEventDate
}

func (b *Bot) handleAddEventStartTime(update *echotron.Update) stateFn {
	b.event.StartTime = update.Message.Text

	_, err := b.SendMessage("Введи время окончания тренировки", b.chatID, nil)
	if err != nil {
		return b.handleMessage
	}

	return b.handleAddEventDate
}

func (b *Bot) handleAddEventEndTime(update *echotron.Update) stateFn {
	b.event.EndTime = update.Message.Text

	_, err := b.SendMessage("Введи стоимость тренировки", b.chatID, nil)
	if err != nil {
		return b.handleMessage
	}

	return b.handleAddEventDate
}

func (b *Bot) handleAddEventPrice(update *echotron.Update) stateFn {
	b.event.Price = update.Message.Text

	_, err := b.SendMessage("Введи максимальное количество участников тренировки", b.chatID, nil)
	if err != nil {
		return b.handleMessage
	}

	return b.handleAddEventDate
}

func (b *Bot) handleAddEventLimit(update *echotron.Update) stateFn {
	b.event.Limit = update.Message.Text

	_, err := b.SendMessage("Введи описание тренировки", b.chatID, nil)
	if err != nil {
		return b.handleMessage
	}

	return b.handleAddEventDate
}

func (b *Bot) handleAddEventDescription(update *echotron.Update) stateFn {
	b.event.Description = update.Message.Text

	id := b.eventClient.AddEvent(b.event.DisciplineId, b.event.Place, b.event.Date, b.event.StartTime, b.event.EndTime, b.event.Price, b.event.Limit, b.event.Description)
	users := b.connectionClient.GetSet("discipline", "users", b.event.DisciplineId)
	for _, user := range users {
		userId, _ := strconv.ParseInt(user, 10, 64)
		_, err := b.SendMessage("Новая тренировка в дисциплине "+b.getDisciplineName(b.event.DisciplineId)+"\nМесто: "+b.event.Place+"\nДата: "+b.event.Date+"\nВремя: "+b.event.StartTime+"-"+b.event.EndTime+"\nСтоимость: "+b.event.Price+"\nОписание: "+b.event.Description, userId, &echotron.MessageOptions{
			ReplyMarkup: eventInlineKeyboard(id, false),
		})

		if err != nil {
			return b.handleMessage
		}
	}

	_, err := b.SendMessage("Тренировка добавлена", b.chatID, &echotron.MessageOptions{
		ReplyMarkup: defaultKeyboard(b.user.IsAdmin),
	})
	if err != nil {
		return b.handleMessage
	}

	return b.handleMessage
}
