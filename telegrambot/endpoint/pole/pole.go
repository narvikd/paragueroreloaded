package pole

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/timekit"
	"paraguero_reloaded/telegrambot"
	"paraguero_reloaded/telegrambot/handler"
	"strings"
)

var countPole = 0

const maxPoles = 3

func Pole(bot *tb.Bot, src *tb.Message) {
	cleanedReceivedMessage := strings.ToLower(src.Text)
	if timekit.IsMidnight() && cleanedReceivedMessage == "pole" && !isPoleExhausted() {
		chatID := tb.ChatID(src.Chat.ID)
		countPole++
		telegrambot.SendMessage(bot, chatID, handleMedal(src))
	}
}

func isPoleExhausted() bool {
	return countPole >= maxPoles
}

func handleMedal(src *tb.Message) string {
	switch countPole {
	case 1:
		return handler.MakeNewMention(src) + " ha hecho la pole. Medalla de oro!"
	case 2:
		return handler.MakeNewMention(src) + " ha hecho la subpole. Medalla de plata!"
	case 3:
		return handler.MakeNewMention(src) + " ha hecho un fail. Pa tu casa, champion"
	default:
		return "Error en la pole"
	}
}
