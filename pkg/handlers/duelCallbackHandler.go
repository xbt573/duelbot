package handlers

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xbt573/duelbot/pkg/phrases"
	"github.com/xbt573/duelbot/pkg/types"
)

var (
	admin, adminIsSet = os.LookupEnv("ADMIN_NAME")
)

func NewDuelCallbackHandler(bot *tgbotapi.BotAPI) types.Handler[tgbotapi.Update] {
	return types.Handler[tgbotapi.Update]{
		Handler: func(update tgbotapi.Update) bool {
			if update.CallbackQuery == nil {
				return false
			}

			if !strings.HasPrefix(update.CallbackQuery.Data, "#duelbot:") {
				return false
			}

			ids := strings.Split(strings.ReplaceAll(update.CallbackData(), "#duelbot:", ""), "-")
			if len(ids) < 2 || len(ids) > 2 {
				return false
			}

			firstId, err := strconv.ParseInt(ids[0], 10, 64)
			if err != nil {
				return false
			}

			secondId, err := strconv.ParseInt(ids[1], 10, 64)
			if err != nil {
				return false
			}

			if secondId != update.CallbackQuery.From.ID {
				response := tgbotapi.NewCallbackWithAlert(update.CallbackQuery.ID, "❌ Access denied!")
				if _, err := bot.Request(response); err != nil {
					panic(err)
				}

				return false
			}

			rand.Seed(time.Now().Unix())
			randomNum := rand.Intn(101)

			var winnerId int64
			var loserId int64

			if randomNum < 50 {
				winnerId = firstId
				loserId = secondId
			} else {
				winnerId = secondId
				loserId = firstId
			}

			winnerUser, err := bot.GetChatMember(tgbotapi.GetChatMemberConfig{
				ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
					ChatID: update.CallbackQuery.Message.Chat.ID,
					UserID: winnerId,
				},
			})

			if err != nil {
				return false
			}

			loserUser, err := bot.GetChatMember(tgbotapi.GetChatMemberConfig{
				ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
					ChatID: update.CallbackQuery.Message.Chat.ID,
					UserID: loserId,
				},
			})

			if err != nil {
				return false
			}

			winnerName := strings.Join([]string{winnerUser.User.FirstName, winnerUser.User.LastName}, " ")
			loserName := strings.Join([]string{loserUser.User.FirstName, loserUser.User.LastName}, " ")

			messageText := fmt.Sprintf("🤠 %v wins! 🤠 %v loses! ⚔️", winnerName, loserName)

			edit := tgbotapi.NewEditMessageTextAndMarkup(
				update.CallbackQuery.Message.Chat.ID,
				update.CallbackQuery.Message.MessageID,
				"⚔️"+phrases.Random(),
				tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{}),
			)

			if _, err := bot.Request(edit); err != nil {
				panic(err)
			}

			time.Sleep(time.Duration(int64(time.Second) * int64(1+rand.Int63n(3))))

			edit = tgbotapi.NewEditMessageTextAndMarkup(
				update.CallbackQuery.Message.Chat.ID,
				update.CallbackQuery.Message.MessageID,
				messageText,
				tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{}),
			)

			if _, err := bot.Request(edit); err != nil {
				panic(err)
			}

			if adminIsSet {
				var loser string

				if loserUser.User.UserName != "" {
					loser = "@" + loserUser.User.UserName
				} else {
					loser = loserName
				}

				msg := tgbotapi.NewMessage(
					update.CallbackQuery.Message.Chat.ID,
					fmt.Sprintf("%v, забань этого мудилу (%v) :D", admin, loser),
				)

				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			}

			return true
		},
	}
}
