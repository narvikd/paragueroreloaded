package endpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/stringkit"
	"paraguero_reloaded/telegrambot"
	"paraguero_reloaded/telegrambot/middlewares"
	"strconv"
	"strings"
)

func handleSendMsgAdminDAW(bot *tb.Bot, route string) {
	bot.Handle(route, func(src *tb.Message) {
		if middlewares.IsAdmin(bot, src) {
			const dawID = -1001155539490
			const offtopicID = -1001361948537

			const chatID = tb.ChatID(dawID)

			prefix := route + " "
			msg := strings.Split(src.Text, prefix)
			telegrambot.SendMessage(bot, chatID, msg[1])
		}
	})
}

func handleSendGroupMsg(bot *tb.Bot, route string) {
	bot.Handle(route, func(src *tb.Message) {
		if middlewares.IsAdmin(bot, src) {
			prefix := route + " "
			msg := strings.Split(src.Text, prefix)
			msg = stringkit.SplitStringNumTimes(msg[1], 14) // Since group IDs in telegram have 14 characters
			chatIDInt, err := strconv.Atoi(msg[0])
			if err != nil {
				telegrambot.SendMessage(bot, tb.ChatID(src.Sender.ID),
					"Error:\n"+
						"Has enviado algo que no es un número como segundo parámetro\n\n"+
						err.Error())
			}
			chatID := tb.ChatID(chatIDInt)
			telegrambot.SendMessage(bot, chatID, msg[1])
		}
	})
}
