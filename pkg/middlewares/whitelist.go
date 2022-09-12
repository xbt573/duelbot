package middlewares

import (
	"gopkg.in/telebot.v3"
)

func Whitelist(chatId int64) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func (ctx telebot.Context) error {
			if ctx.Chat().ID == chatId {
				return next(ctx)
			}

			return nil
		}
	}
}
