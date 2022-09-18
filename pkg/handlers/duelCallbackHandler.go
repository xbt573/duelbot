package handlers

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xbt573/duelbot/pkg/phrases"
	"github.com/xbt573/duelbot/pkg/types"
	"gopkg.in/telebot.v3"
)

func DuelCallbackHandler(ctx telebot.Context) error {
	data := strings.ReplaceAll(ctx.Callback().Data, "\fduel|", "")
	ids := strings.Split(strings.ReplaceAll(data, "#duelbot:", ""), "-")
	if len(ids) < 2 || len(ids) > 2 {
		return ctx.Respond()
	}

	firstId, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil {
		return ctx.Respond()
	}

	secondId, err := strconv.ParseInt(ids[1], 10, 64)
	if err != nil {
		return ctx.Respond()
	}

	if secondId != ctx.Callback().Sender.ID {
		return ctx.Respond(&telebot.CallbackResponse{Text: "❌ Access denied!"})
	}

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(2)

	var winnerId int64
	var loserId int64

	if randomNum == 1 {
		winnerId = firstId
		loserId = secondId
	} else {
		winnerId = secondId
		loserId = firstId
	}

	winnerUser, err := ctx.Bot().ChatMemberOf(ctx.Chat(), &types.IdRecipient{Id: winnerId})
	if err != nil {
		return ctx.Respond()
	}

	loserUser, err := ctx.Bot().ChatMemberOf(ctx.Chat(), &types.IdRecipient{Id: loserId})
	if err != nil {
		return ctx.Respond()
	}

	winnerName := strings.Join([]string{winnerUser.User.FirstName, winnerUser.User.LastName}, " ")
	loserName := strings.Join([]string{loserUser.User.FirstName, loserUser.User.LastName}, " ")

	messageText := fmt.Sprintf("🤠 %v wins! 🤠 %v loses! ⚔️", winnerName, loserName)

	_, err = ctx.Bot().Edit(ctx.Callback(), "⚔️"+phrases.Random(), &telebot.ReplyMarkup{})
	if err != nil {
		return ctx.Respond()
	}

	time.Sleep(time.Duration(int64(time.Second) * int64(1+rand.Int63n(3))))

	_, err = ctx.Bot().Edit(ctx.Callback(), messageText)
	if err != nil {
		return ctx.Respond()
	}

	mode, modeIsSet := os.LookupEnv("BOT_MODE")
	if !modeIsSet {
		return ctx.Respond()
	}

	switch mode {
	case "PING":
		admin, adminIsSet := os.LookupEnv("ADMIN_NAME")

		if !adminIsSet {
			return ctx.Respond()
		}

		var loser string

		if loserUser.User.Username != "" {
			loser = "@" + loserUser.User.Username
		} else {
			loser = loserName
		}

		err := ctx.Send(fmt.Sprintf("%v, ban him (%v)! :D", admin, loser))
		if err != nil {
			return ctx.Respond()
		}

	case "ADMIN":
		me, err := ctx.Bot().ChatMemberOf(ctx.Chat(), ctx.Bot().Me)
		if err != nil {
			return ctx.Respond()
		}

		if !me.CanRestrictMembers {
			return ctx.Reply("Bot is not an admin. Change mode or give rights to restrict members!")
		}

		loserUser.RestrictedUntil = time.Now().Add(time.Minute * 10).Unix()
		loserUser.CanSendMessages = false

		err = ctx.Bot().Restrict(ctx.Chat(), loserUser)
		if err != nil {
			return ctx.Respond()
		}
	}

	return ctx.Respond()
}
