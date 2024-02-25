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

type Subscription struct {
	Channels []int64 `json:"channels"`
	Name     string  `json:"name"`
}

var (
	bot       *TelegramBot.BotAPI
	err       error
	now       = time.Now()
	parser    = Feed.NewParser()
	updatedAt time.Time
)

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

func getUpdatedAt() {
	bytes, err := os.ReadFile("/suppress/state/updatedAt")
	if err != nil {
		log.Println("Could not read time: ", err)
		updatedAt = now
		return
	}
	err = json.Unmarshal(bytes, &updatedAt)
	if err != nil {
		log.Println("Could not deserialize time: ", err)
	}
}

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

	getUpdatedAt()

	subscriptions := subscriptions()

	var group sync.WaitGroup
	group.Add(len(subscriptions))

	for endpoint, subscription := range subscriptions {
		go fetch(&group, endpoint, subscription.Channels)
	}

	group.Wait()

	setUpdatedAt()
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

func setUpdatedAt() {
	bytes, err := json.Marshal(now)
	if err != nil {
		log.Fatal("Could not serialize time: ", err)
	}

	err = os.WriteFile("/suppress/state/updatedAt", bytes, 0400)
	if err != nil {
		log.Fatal("Could not write time: ", err)
	}
}

func subscriptions() map[string]Subscription {
	path := "/suppress/configuration/subscriptions.json"
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Could not read subscriptions: ", err)
	}

	var subscriptions map[string]Subscription
	err = json.Unmarshal(file, &subscriptions)
	if err != nil {
		log.Fatal("Could not deserialize subscriptions: ", err)
	}

	return subscriptions
}
