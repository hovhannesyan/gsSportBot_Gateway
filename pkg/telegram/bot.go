package telegram

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/hovhannesyan/gsSportBot_Gateway/pkg/connection"
	"github.com/hovhannesyan/gsSportBot_Gateway/pkg/event"
	"github.com/hovhannesyan/gsSportBot_Gateway/pkg/models"
	"github.com/hovhannesyan/gsSportBot_Gateway/pkg/user"
)

type stateFn func(*echotron.Update) stateFn

type Bot struct {
	chatID int64
	state  stateFn
	user   *models.User
	event  *models.Event
	echotron.API
	userClient       *user.ServiceClient
	eventClient      *event.ServiceClient
	connectionClient *connection.ServiceClient
}

func NewBot(token, userServiceURL, eventServiceURL, connectionServiceURL string) func(chatID int64) echotron.Bot {
	return func(chatID int64) echotron.Bot {
		bot := &Bot{
			chatID:           chatID,
			API:              echotron.NewAPI(token),
			userClient:       user.InitServiceClient(userServiceURL),
			eventClient:      event.InitServiceClient(eventServiceURL),
			connectionClient: connection.InitServiceClient(connectionServiceURL),
		}

		bot.state = bot.handleMessage
		return bot
	}
}

func (b *Bot) Update(update *echotron.Update) {
	b.state = b.state(update)
}

func (b *Bot) handleMessage(update *echotron.Update) stateFn {
	if update.Message.Text == "/start" {
		return b.handleStart(update)
	}
	if b.user == nil {
		data := b.userClient.GetUsersById(update.Message.From.ID)
		if data == nil {
			_, err := b.SendMessage("Ты не зареган, для регистрации используй команду\n/start", b.chatID, nil)
			if err != nil {
				return b.handleMessage
			}
			return b.handleMessage
		}
		b.user = &models.User{
			Id:       data[0].Id,
			Username: data[0].Username,
			Name:     data[0].Name,
			Phone:    data[0].Phone,
			IsAdmin:  data[0].IsAdmin,
		}
		_, err := b.SendMessage("Я тебя	авторизовал, можешь пользоваться ботом как раньше", b.chatID, &echotron.MessageOptions{
			ReplyMarkup: defaultKeyboard(b.user.IsAdmin),
		})
		if err != nil {
			return b.handleMessage
		}
	}
	if update.Message.Text == "Настройки" {
		return b.handleSettings(update)
	}
	if update.Message.Text == "Добавить дисциплину" {
		return b.handleAddDiscipline(update)
	}
	return b.handleMessage
}
