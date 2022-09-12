package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/xbt573/duelbot/pkg/handlers"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func main() {
	bot, err := telebot.NewBot(telebot.Settings{
		Token: os.Getenv("BOT_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initialized bot")

	chatId, chatIdIsSet := os.LookupEnv("CHAT_ID")
	if !chatIdIsSet {
		log.Fatal("Chat ID is not set!")
	}

	chatIdParsed, err := strconv.ParseInt(chatId, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	bot.Use(middleware.Whitelist(chatIdParsed))

	handlers.Route(bot)
	log.Println("Initialized routes")

	log.Println("Started!")
	bot.Start()
}
