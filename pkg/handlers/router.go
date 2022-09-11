package handlers

import (
	"gopkg.in/telebot.v3"
)

func Route(bot *telebot.Bot) {
	bot.Handle("/start", StartHandler)
	bot.Handle("/help", HelpHandler)
	bot.Handle("/duel", DuelHandler)
	bot.Handle(telebot.OnCallback, DuelCallbackHandler)
}
