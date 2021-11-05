package onuserjoin

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot/handler"
)

var funcSlice []func(bot *tb.Bot, src *tb.Message)

func loadEndpoints(bot *tb.Bot) {
	bot.Handle(tb.OnUserJoined, func(src *tb.Message) {
		handler.IterateEndpoints(bot, src, funcSlice)
	})
}

func AddEndpoint(bot *tb.Bot, function func(bot *tb.Bot, src *tb.Message)) {
	funcSlice = append(funcSlice, function)
	loadEndpoints(bot)
}
