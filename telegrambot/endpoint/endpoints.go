package endpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot/endpoint/andalu"
	"paraguero_reloaded/telegrambot/endpoint/debug"
	"paraguero_reloaded/telegrambot/endpoint/pole"
	"paraguero_reloaded/telegrambot/endpoint/rust"
	"paraguero_reloaded/telegrambot/endpoint/kotlin"
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
	handler.MsgEndpoint(bot, "/java", getJavaCourses())
	onuserjoin.AddEndpoint(bot, welcomeendpoint.Welcome)
	ontext.AddEndpoint(bot, paraguas)
	ontext.AddEndpoint(bot, pole.Pole)
	ontext.AddEndpoint(bot, rust.GetRustMotivation)
	ontext.AddEndpoint(bot, kotlin.GetKotlinMotivation)
	ontext.AddEndpoint(bot, andalu.TranslateToAndalu)
	ontext.AddEndpoint(bot, translateToChoni)
	onsticker.AddEndpoint(bot, paraguasSticker)
	handleSendMsgAdminDAW(bot, "/daw")
	handleSendGroupMsg(bot, "/group")
	debugendpoint.GetCurrentTime(bot, "/time")
	debugendpoint.GetChatID(bot, "/chatid")

	// UNUSED:
	// debugEndpoint.PrintSrc(bot, "/src")
	//debugEndpoint.GetUsername(bot)
	//onsticker.AddEndpoint(bot, getStickerInfo)
	//welcomeEndpoint.Debug(bot, "!welcome")
	//ontext.AddEndpoint(bot, debugEndpoint.GetSenderID)
}
