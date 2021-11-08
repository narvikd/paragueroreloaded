package welcomeendpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/logger"
	"paraguero_reloaded/telegrambot/handler"
)

func Welcome(bot *tb.Bot, src *tb.Message) {
	chatID := tb.ChatID(src.Chat.ID)
	senderMention := handler.MakeNewMention(src)
	_, err := bot.Send(chatID, getWelcomeMsg(senderMention), "html")
	logger.LogIfError(err, "error")
}

func Debug(bot *tb.Bot, route string) {
	bot.Handle(route, func(src *tb.Message) {
		Welcome(bot, src)
	})
}
