package welcomeEndpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/logger"
	"strconv"
)

func Welcome(bot *tb.Bot, src *tb.Message) {
	chatID := tb.ChatID(src.Chat.ID)
	userName := src.Sender.Username
	senderMention := "<a href=\"tg://user?id=" + strconv.Itoa(src.Sender.ID) + "\">@" + userName + "</a>"
	if userName == "" {
		senderMention = "<a href=\"tg://user?id=" + strconv.Itoa(src.Sender.ID) + "\">" + src.Sender.FirstName + "</a>"
	}
	_, err := bot.Send(chatID, getMsg(senderMention), "html")
	logger.LogIfError(err, "error")
}

func Debug(bot *tb.Bot, route string) {
	bot.Handle(route, func(src *tb.Message) {
		Welcome(bot, src)
	})
}
