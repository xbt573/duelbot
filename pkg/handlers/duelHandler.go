package handlers

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xbt573/duelbot/pkg/types"
)

func NewDuelHandler(bot *tgbotapi.BotAPI) types.Handler[tgbotapi.Update] {
	return types.Handler[tgbotapi.Update]{
		Handler: func(update tgbotapi.Update) bool {
			if update.Message == nil {
				return false
			}

			if !update.Message.IsCommand() {
				return false
			}

			if update.Message.Command() != "duel" {
				return false
			}

			if update.Message.ReplyToMessage == nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Command should be reply to message!")
				msg.ReplyToMessageID = update.Message.MessageID

				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			}

			first, err := bot.GetChatMember(tgbotapi.GetChatMemberConfig{
				ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
					ChatID: update.Message.Chat.ID,
					UserID: update.Message.From.ID,
				},
			})

			if err != nil {
				panic(err)
			}

			second, err := bot.GetChatMember(tgbotapi.GetChatMemberConfig{
				ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
					ChatID: update.Message.Chat.ID,
					UserID: update.Message.ReplyToMessage.From.ID,
				},
			})

			firstName := strings.Join([]string{ first.User.FirstName, first.User.LastName }, " ")
			secondName := strings.Join([]string{ second.User.FirstName, second.User.LastName }, " ")

			messageText := fmt.Sprintf("🤠 %v challenges 🤠 %v to a duel! 🔫", firstName, secondName)
			callbackData := fmt.Sprintf("#duelbot:%v-%v", first.User.ID, second.User.ID)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("✅ Accept", callbackData),
				),
			)

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}

			return true
		},
	}
}
