package main

import (
	"log"
	"os"
	"time"

	"github.com/xbt573/duelbot/pkg/handlers"
	"github.com/xbt573/duelbot/pkg/tdapi"
	"gopkg.in/telebot.v3"
)

func main() {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("BOT_MODE") == "USER" {
		client, err := tdapi.New()
		if err != nil {
			log.Fatal(err)
		}

		bot.Use(func(next telebot.HandlerFunc) telebot.HandlerFunc {
			return func(ctx telebot.Context) error {
				ctx.Set("client", client)
				return next(ctx)
			}
		})
	}

	log.Println("Initialized bot")

	handlers.Route(bot)
	log.Println("Initialized routes")

	log.Println("Started!")
	bot.Start()
}
