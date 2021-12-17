package endpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot/endpoint/api/randmsgapi"
	"paraguero_reloaded/telegrambot/endpoint/pole"
	"paraguero_reloaded/telegrambot/endpoint/stats"
	"paraguero_reloaded/telegrambot/endpoint/welcome"
	"paraguero_reloaded/telegrambot/handler/onsticker"
	"paraguero_reloaded/telegrambot/handler/ontext"
	"paraguero_reloaded/telegrambot/handler/onuserjoin"
)

// LoadEndpoints loads all the endpoints, and it's where they're declared
func LoadEndpoints(bot *tb.Bot) {
	// Users
	randmsgapi.Get(bot, "/help")
	randmsgapi.Get(bot, "/java")
	ontext.AddEndpoint(bot, stats.SendStats)
	links(bot, "/links")
	ontext.AddEndpoint(bot, paraguas)
	onsticker.AddEndpoint(bot, paraguasSticker)
	ontext.AddEndpoint(bot, pole.Run)
	onuserjoin.AddEndpoint(bot, welcome.Run)
	// User jokes
	randmsgapi.Get(bot, "/rust")
	randmsgapi.Get(bot, "/kotlin")
}
