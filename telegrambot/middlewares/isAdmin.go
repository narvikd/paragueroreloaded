package middlewares

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot"
)

func IsAdmin(bot *tb.Bot, src *tb.Message) bool {
	const adminID = 1099020633
	if src.Sender.ID == adminID {
		return true
	}
	chatID := tb.ChatID(src.Chat.ID)
	telegrambot.SendMessage(bot, chatID, "Venga tonto, pa tu casa")
	return false
}
