package endpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	debugEndpoint "paraguero_reloaded/telegrambot/endpoint/debug"
	"paraguero_reloaded/telegrambot/endpoint/pole"
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
	onuserjoin.AddEndpoint(bot, welcomeEndpoint.Welcome)
	ontext.AddEndpoint(bot, paraguas)
	ontext.AddEndpoint(bot, pole.Pole)
	onsticker.AddEndpoint(bot, paraguasSticker)
	handleSendMsgAdminDAW(bot, "/daw")
	handleSendGroupMsg(bot, "/group")
	debugEndpoint.GetCurrentTime(bot, "/time")
	debugEndpoint.GetChatID(bot, "/chatid")

	// UNUSED:
	// debugEndpoint.PrintSrc(bot, "/src")
	//debugEndpoint.GetUsername(bot)
	//onsticker.AddEndpoint(bot, getStickerInfo)
	//welcomeEndpoint.Debug(bot, "!welcome")
	//ontext.AddEndpoint(bot, debugEndpoint.GetSenderID)
}
