package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xbt573/duelbot/pkg/handlers"
	"github.com/xbt573/duelbot/pkg/types"
)

var (
	token, tokenIsSet = os.LookupEnv("BOT_TOKEN")
)

func initHandlerChain(bot *tgbotapi.BotAPI) types.Handler[tgbotapi.Update] {
	startHandler := handlers.NewStartHandler(bot)
	helpHandler := handlers.NewHelpHandler(bot)
	duelHandler := handlers.NewDuelHandler(bot)
	duelCallbackHandler := handlers.NewDuelCallbackHandler(bot)

	startHandler.SetNext(&helpHandler)
	helpHandler.SetNext(&duelHandler)
	duelHandler.SetNext(&duelCallbackHandler)

	return startHandler
}

func handleUpdate(update tgbotapi.Update, chain types.Handler[tgbotapi.Update]) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovered update %v! Error: %v\n", update.UpdateID, err)
		}
	}()

	chain.Handle(update)
}

func main() {
	if !tokenIsSet {
		log.Panic("Token is not set!")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized as %v\n", bot.Self.UserName)

	chain := initHandlerChain(bot)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		go handleUpdate(update, chain)
	}
}
