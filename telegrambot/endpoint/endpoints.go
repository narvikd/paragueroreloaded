package endpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot/endpoint/welcome"
	"paraguero_reloaded/telegrambot/handler"
	"paraguero_reloaded/telegrambot/handler/onsticker"
	"paraguero_reloaded/telegrambot/handler/ontext"
	"paraguero_reloaded/telegrambot/handler/onuserjoin"
)

// LoadEndpoints loads all the endpoints, and it's where they're declared
func LoadEndpoints(bot *tb.Bot) {
	links(bot, "/links")
	handler.MsgEndpoint(bot, "/ban", "Venga tonto, pa tu casa")
	onuserjoin.AddEndpoint(bot, welcomeEndpoint.Welcome)
	ontext.AddEndpoint(bot, paraguas)
	onsticker.AddEndpoint(bot, paraguasSticker)

	// UNUSED:
	//debug.GetUsername(bot)
	//onsticker.AddEndpoint(bot, getStickerInfo)
	//welcomeEndpoint.Debug(bot, "!welcome")
}
