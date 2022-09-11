package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xbt573/duelbot/pkg/types"
)

func NewHelpHandler(bot *tgbotapi.BotAPI) types.Handler[tgbotapi.Update] {
	return types.Handler[tgbotapi.Update]{
		Handler: func(update tgbotapi.Update) bool {
			if update.Message == nil {
				return false
			}

			if !update.Message.IsCommand() {
				return false
			}

			if update.Message.Command() != "help" {
				return false
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, `
					In order to duel send /duel command in reply to message
				`)
			msg.ReplyToMessageID = update.Message.MessageID

			_, err := bot.Send(msg)
			if err != nil {
				panic(err)
			}

			return true
		},
	}
}
