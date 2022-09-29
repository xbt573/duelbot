package handlers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/telebot.v3"
)

func SuicideHandler(ctx telebot.Context) error {
	mode, modeIsSet := os.LookupEnv("BOT_MODE")
	if !modeIsSet {
		return ctx.Reply("Bot mode is not set. Contact system administrator")
	}

	user, err := ctx.Bot().ChatMemberOf(ctx.Chat(), ctx.Sender())
	if err != nil {
		return err
	}

	var username string

	if user.User.Username != "" {
		username = user.User.Username
	} else {
		username = strings.Join([]string{user.User.FirstName, user.User.LastName}, " ")
	}

	switch mode {
	case "PING":
		admin, adminIsSet := os.LookupEnv("ADMIN_NAME")

		if !adminIsSet {
			return ctx.Reply("Admin name is not set. Contact system administrator")
		}

		return ctx.Reply(
			fmt.Sprintf("%v, mute him (%v) for 20 minutes!", admin, username),
			admin,
			username,
		)

	case "ADMIN":
		me, err := ctx.Bot().ChatMemberOf(ctx.Chat(), ctx.Bot().Me)
		if err != nil {
			return err
		}

		if !me.CanRestrictMembers {
			return ctx.Reply("Bot is not an admin. Contact chat admin")
		}

		user.CanSendMessages = false
		user.RestrictedUntil = time.Now().Add(time.Minute * 20).Unix()

		return ctx.Bot().Restrict(ctx.Chat(), user)
	}

	return nil
}
