package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	TelegramBot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	Feed "github.com/mmcdole/gofeed"
)

var (
	bot       *TelegramBot.BotAPI
	err       error
	parser    = Feed.NewParser()
	updatedAt = time.Now()
)

func main() {
	token, exists := os.LookupEnv("TELEGRAM_BOT_TOKEN")
	if !exists {
		log.Fatal("TELEGRAM_BOT_TOKEN not set.")
	}

	// 1. Get the chat ID: `https://api.telegram.org/bot${token}/getUpdates`
	// 2. Send a message: `https://api.telegram.org/bot${token}/sendMessage?chat_id=${chat}&text=Hi!`

	bot, err = TelegramBot.NewBotAPI(token)
	if err != nil {
		log.Panic("Could not create Telegram bot: ", err)
	}

	go func() {
		for {
			subscriptions := subscriptions()

			var group sync.WaitGroup
			group.Add(len(subscriptions))

			for subscription, chats := range subscriptions {
				go fetch(&group, subscription, chats)
			}

			group.Wait()
			updatedAt = time.Now()

			time.Sleep(30 * time.Minute)
		}
	}()

	select {}
}

func fetch(group *sync.WaitGroup, subscription string, chats []int64) {
	defer group.Done()
	feed, err := parser.ParseURL(subscription)
	if err != nil {
		log.Printf("Could not fetch RSS feed %s: %v\n", subscription, err)
		return
	}

	if feed.Items == nil {
		return
	}

	for _, item := range feed.Items {
		publishedAt := item.PublishedParsed
		if publishedAt == nil {
			publishedAt = item.UpdatedParsed
		}

		if publishedAt.After(updatedAt) {
			var group sync.WaitGroup
			group.Add(len(chats))
			for _, chat := range chats {
				go send(&group, bot, chat, item.Link)
			}
		}
	}
}

func send(group *sync.WaitGroup, bot *TelegramBot.BotAPI, chat int64, link string) {
	defer group.Done()
	message := TelegramBot.NewMessage(chat, link)
	message.ParseMode = "Markdown"
	_, err := bot.Send(message)
	if err != nil {
		log.Printf("Could not send message to Telegram chat %d: %v\n", chat, err)
	}
}

func subscriptions() map[string][]int64 {
	path := "/suppress/subscriptions.json"
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Could not read subscriptions: ", err)
	}

	var subscriptions map[string][]int64
	err = json.Unmarshal(file, &subscriptions)
	if err != nil {
		log.Fatal("Could not parse subscriptions: ", err)
	}

	return subscriptions
}
