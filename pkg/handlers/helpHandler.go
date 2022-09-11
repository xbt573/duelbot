package handlers

import (
	"gopkg.in/telebot.v3"
)

func HelpHandler(ctx telebot.Context) error {
	return ctx.Reply(`
		In order to duel send /duel command in reply to message
	`, &telebot.SendOptions{
		ReplyTo: ctx.Message(),
	})
}
