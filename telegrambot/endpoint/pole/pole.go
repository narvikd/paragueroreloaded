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
		msg := handler.MakeNewMention(src) + " " + getPoleMedal()
		telegrambot.SendMessage(bot, chatID, msg)
	}
}

func isPoleExhausted() bool {
	return countPole >= maxPoles
}

func getPoleMedal() string {
	switch countPole {
	case 1:
		return "ha hecho la pole. Medalla de oro!"
	case 2:
		return "ha hecho la subpole. Medalla de plata!"
	case 3:
		return "ha hecho un fail. Pa tu casa, champion"
	default:
		return "Error en la pole"
	}
}
