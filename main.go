package main

import (
	"log"
	"os"
	"time"

	"gopkg.in/telebot.v3"
	"github.com/xbt573/duelbot/pkg/handlers"
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

	handlers.Route(bot)
	log.Println("Initialized routes")

	log.Println("Started!")
	bot.Start()
}
