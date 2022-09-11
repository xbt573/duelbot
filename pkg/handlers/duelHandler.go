package handlers

import (
	"fmt"
	"strings"

	"gopkg.in/telebot.v3"
)

func DuelHandler(ctx telebot.Context) error {
	if ctx.Message().ReplyTo == nil {
		return ctx.Reply("Command should be a reply to message!", &telebot.SendOptions{
			ReplyTo: ctx.Message(),
		})
	}

	if ctx.Message().ReplyTo.Sender == ctx.Sender() {
		return ctx.Reply("You can't duel yourself!", &telebot.SendOptions{
			ReplyTo: ctx.Message(),
		})
	}

	if ctx.Message().ReplyTo.Sender == ctx.Bot().Me {
		return ctx.Reply("You can't duel bot!", &telebot.SendOptions{
			ReplyTo: ctx.Message(),
		})
	}

	firstName := strings.Join([]string{ctx.Sender().FirstName, ctx.Sender().LastName}, " ")
	secondName := strings.Join([]string{ctx.Message().ReplyTo.Sender.FirstName, ctx.Message().ReplyTo.Sender.LastName}, " ")

	messageText := fmt.Sprintf("🤠 %v challenges 🤠 %v to a duel! 🔫", firstName, secondName)
	callbackData := fmt.Sprintf("#duelbot:%v-%v", ctx.Sender().ID, ctx.Message().ReplyTo.Sender.ID)

	markup := &telebot.ReplyMarkup{}
	markup.Inline(
		markup.Row(
			markup.Data("✅ Accept", "duel", callbackData),
		),
	)

	return ctx.Reply(messageText, &telebot.SendOptions{
		ReplyTo: ctx.Message().ReplyTo,
		ReplyMarkup: markup,
	})
}
