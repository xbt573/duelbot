package handlers

import (
	"log"
	"os"
	"strconv"

	"github.com/xbt573/duelbot/pkg/middlewares"
	"gopkg.in/telebot.v3"
)

func Route(bot *telebot.Bot) {
	chatId, chatIdIsSet := os.LookupEnv("CHAT_ID")
	if !chatIdIsSet {
		log.Fatal("Chat ID is not set!")
	}

	chatIdParsed, err := strconv.ParseInt(chatId, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	bot.Use(middlewares.Whitelist(chatIdParsed))

	bot.Handle("/start", StartHandler)
	bot.Handle("/help", HelpHandler)
	bot.Handle("/duel", DuelHandler)
	bot.Handle("/suicide", SuicideHandler)
	bot.Handle("/rm", RemoveHandler)
	bot.Handle(telebot.OnCallback, DuelCallbackHandler)
}
