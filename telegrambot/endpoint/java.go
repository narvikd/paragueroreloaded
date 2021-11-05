package endpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/stringkit"
	"paraguero_reloaded/telegrambot"
	"strings"
)

func java(bot *tb.Bot, src *tb.Message) {
	const pattern = "\\bjava\\b"
	cleanedReceivedMessage := strings.ToLower(src.Text)

	if stringkit.IsRegexMatch(cleanedReceivedMessage, pattern) {
		chatID := tb.ChatID(src.Chat.ID)
		const msg = "Decir que java es bueno porque funciona en todos lados es como decir que el sexo anal " +
			"es bueno porque funciona en todos los géneros."

		telegrambot.SendMessage(bot, chatID, msg)
	}
}
