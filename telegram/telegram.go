package telegram

import (
	"github.com/tucnak/telebot"
	"time"
)

type Telegram struct {
	bot    *telebot.Bot
	chatID int64
}

// NewSender создает и настраивает нового бота для отправки уведомлений.
func NewSender(token string, chatID int64) (*Telegram, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}

	return &Telegram{
		bot:    bot,
		chatID: chatID,
	}, nil
}

// SendNotification отправляет уведомление в Telegram.
func (t *Telegram) SendNotification(message string) error {
	chat := &telebot.Chat{ID: t.chatID} // Создаем Chat с использованием ID
	_, err := t.bot.Send(chat, message)
	return err
}
