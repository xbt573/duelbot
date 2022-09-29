package handlers

import "gopkg.in/telebot.v3"

func RemoveHandler(ctx telebot.Context) error {
	if ctx.Message().ReplyTo.Sender.ID != ctx.Bot().Me.ID {
		return ctx.Reply("You can delete only messages of this bot")
	}

	if ctx.Message().ReplyTo.ReplyMarkup != nil {
		return ctx.Reply("Duel is pending, you can't delete duel messages")
	}

	return ctx.Bot().Delete(ctx.Message().ReplyTo)
}
