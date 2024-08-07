package notifier

import (
	"context"
	"time"

	"github.com/go-telegram/bot"
	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/logger"
)

type telegram_bot_t struct {
	ctx    context.Context
	cancel context.CancelFunc
	client *bot.Bot
}

var botInstance telegram_bot_t

func InitTelegramNotifier() {
	ctx, cancel := context.WithCancel(context.Background())
	botInstance = telegram_bot_t{
		ctx:    ctx,
		cancel: cancel,
	}

	token := config.Config.Alarm.Telegram.Token
	b, err := bot.New(token)
	if err != nil {
		panic(err)
	}

	_, err = b.DeleteWebhook(botInstance.ctx, &bot.DeleteWebhookParams{})
	if err != nil {
		logger.Log.Error("Failed to delete webhook:", err)
		return
	}

	go b.Start(botInstance.ctx)

	botInstance.client = b
	logger.Log.Info("Telegram bot started successfully.")
}

func TelegramNotifier(message string) {
	msgCtx, cancel := context.WithTimeout(botInstance.ctx, 10*time.Second)
	defer cancel()

	_, err := botInstance.client.SendMessage(msgCtx, &bot.SendMessageParams{
		ChatID: config.Config.Alarm.Telegram.ChatID,
		Text:   message,
	})

	if err != nil {
		logger.Log.Errorf("Failed to send message: %v\n", err)
	}
}
