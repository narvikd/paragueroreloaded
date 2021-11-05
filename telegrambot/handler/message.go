package handler

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot"
)

func MsgEndpoint(bot *tb.Bot, route string, message string) {
	bot.Handle(route, func(src *tb.Message) {
		chatID := tb.ChatID(src.Chat.ID)
		telegrambot.SendMessage(bot, chatID, message)
	})
}
