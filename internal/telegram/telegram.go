package telegram

import (
	"log/slog"
	"os"
	"sync"

	TelegramBot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	Bot *TelegramBot.BotAPI
	err error
)

func Initialize() {
	// 1. Get the chat ID: `https://api.telegram.org/bot${token}/getUpdates`
	// 2. Send a message: `https://api.telegram.org/bot${token}/sendMessage?chat_id=${chat}&text=Hi!`

	token, exists := os.LookupEnv("TELEGRAM_BOT_TOKEN")
	if !exists {
		slog.Error("TELEGRAM_BOT_TOKEN not set.")
		os.Exit(1)
	}

	Bot, err = TelegramBot.NewBotAPI(token)
	if err != nil {
		slog.Error("Could not create Telegram bot.", "error", err)
		os.Exit(1)
	}
}

func Send(group *sync.WaitGroup, chat int64, link string) {
	defer group.Done()
	message := TelegramBot.NewMessage(chat, link)
	message.ParseMode = "Markdown"
	_, err := Bot.Send(message)
	if err != nil {
		slog.Warn("Could not send message to Telegram chat.", "chat", chat, "error", err)
	}
}
