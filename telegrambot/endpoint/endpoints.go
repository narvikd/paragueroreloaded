package endpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot/endpoint/debug"
	"paraguero_reloaded/telegrambot/endpoint/kotlin"
	"paraguero_reloaded/telegrambot/endpoint/pole"
	"paraguero_reloaded/telegrambot/endpoint/rust"
	"paraguero_reloaded/telegrambot/endpoint/welcome"
	"paraguero_reloaded/telegrambot/handler"
	"paraguero_reloaded/telegrambot/handler/onsticker"
	"paraguero_reloaded/telegrambot/handler/ontext"
	"paraguero_reloaded/telegrambot/handler/onuserjoin"
)

// LoadEndpoints loads all the endpoints, and it's where they're declared
func LoadEndpoints(bot *tb.Bot) {
	// Users
	handler.MsgEndpoint(bot, "/help", getHelp())
	links(bot, "/links")
	handler.MsgEndpoint(bot, "/java", getJavaCourses())
	ontext.AddEndpoint(bot, paraguas)
	onsticker.AddEndpoint(bot, paraguasSticker)
	ontext.AddEndpoint(bot, pole.Pole)
	onuserjoin.AddEndpoint(bot, welcomeendpoint.Welcome)
	// User jokes
	handler.MsgEndpoint(bot, "/ban", "Venga tonto, pa tu casa")
	ontext.AddEndpoint(bot, rust.GetRustMotivation)
	ontext.AddEndpoint(bot, kotlin.GetKotlinMotivation)
	// Admins
	handleSendMsgAdminDAW(bot, "/daw")
	handleSendGroupMsg(bot, "/group")
	// Admin Debug
	debugendpoint.GetCurrentTime(bot, "/time")
	debugendpoint.GetChatID(bot, "/chatid")

	// UNUSED:
	// debugEndpoint.PrintSrc(bot, "/src")
	//debugEndpoint.GetUsername(bot)
	//onsticker.AddEndpoint(bot, getStickerInfo)
	//welcomeEndpoint.Debug(bot, "!welcome")
	//ontext.AddEndpoint(bot, debugEndpoint.GetSenderID)
}
