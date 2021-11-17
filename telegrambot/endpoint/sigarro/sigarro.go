package sigarro

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot"
	"paraguero_reloaded/telegrambot/endpoint/andalu"
	"paraguero_reloaded/telegrambot/endpoint/choni"
	"strings"
)

func TranslateToSigarro(bot *tb.Bot, src *tb.Message) {
	if strings.HasPrefix(src.Text, "/sigarro ") {
		chatID := tb.ChatID(src.Chat.ID)

		toTranslate := strings.SplitAfter(src.Text, "/sigarro ")

		_, andaluMsg, errAndalu := andalu.Common(bot, src, toTranslate[1])
		if errAndalu == nil {
			_, choniMsg, _ := choni.Common(bot, src, andaluMsg.Andaluh)
			telegrambot.SendMessage(bot, chatID, choniMsg)
		}
	}
}
