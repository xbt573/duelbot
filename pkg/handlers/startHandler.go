package handlers

import (
	"gopkg.in/telebot.v3"
)

func StartHandler(ctx telebot.Context) error {
	return ctx.Reply(`
		Hello! In order to duel send /duel command in reply to message
	`, &telebot.SendOptions{
		ReplyTo: ctx.Message(),
	})
}
