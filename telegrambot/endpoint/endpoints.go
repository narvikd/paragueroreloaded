package endpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot/endpoint/api/randmsgapi"
	"paraguero_reloaded/telegrambot/endpoint/pole"
	"paraguero_reloaded/telegrambot/endpoint/stats"
	"paraguero_reloaded/telegrambot/endpoint/welcome"
	"paraguero_reloaded/telegrambot/handler/logger"
	"paraguero_reloaded/telegrambot/handler/onsticker"
	"paraguero_reloaded/telegrambot/handler/ontext"
	"paraguero_reloaded/telegrambot/handler/onuserjoin"
	"paraguero_reloaded/telegrambot/handler/whitelist"
)

// LoadEndpoints loads all the endpoints, and it's where they're declared
func LoadEndpoints(bot *tb.Bot) {
	// Logger
	ontext.AddEndpoint(bot, logger.Register)
	// Users
	randmsgapi.Get(bot, "/help")
	randmsgapi.Get(bot, "/java")
	randmsgapi.Get(bot, "/javascript")
	randmsgapi.Get(bot, "/consejos")
	ontext.AddEndpoint(bot, whitelist.Only(stats.SendStats))
	links(bot, "/links")
	ontext.AddEndpoint(bot, whitelist.Only(paraguas))
	onsticker.AddEndpoint(bot, whitelist.Only(paraguasSticker))
	ontext.AddEndpoint(bot, whitelist.Only(pole.Run))
	onuserjoin.AddEndpoint(bot, whitelist.Only(welcome.Run))
	// User jokes
	randmsgapi.Get(bot, "/rust")
	randmsgapi.Get(bot, "/kotlin")
}
