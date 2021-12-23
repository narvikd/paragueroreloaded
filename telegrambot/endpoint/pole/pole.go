package pole

import (
	"github.com/narvikd/timekit"
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/stringkit"
	"paraguero_reloaded/telegrambot"
	"paraguero_reloaded/telegrambot/handler"
	"strconv"
	"strings"
)

const maxPoles = 3

var countPole = 0
var poleadoresIDs []string

// TODO: Refactor this, and instead of receiving "subpole" in "pole" as ok, only accept appropriately strings on appropriately pole positions.
var acceptedPoleStrings = []string{"pole", "mastil", "mástil", "oro",
	"subpole", "plata",
	"fail", "bronce"}

func Run(bot *tb.Bot, src *tb.Message) {
	chatID := tb.ChatID(src.Chat.ID)
	cleanedReceivedMessage := strings.ToLower(src.Text)
	isPoleMsgReceived := stringkit.SliceContains(acceptedPoleStrings, cleanedReceivedMessage)
	if timekit.IsMidnight() && isPoleMsgReceived && !isPoleExhausted() {
		if haveINotSeenThisPoleadorID(src.Sender.ID) {
			telegrambot.SendMessage(bot, chatID, handleMedal(src))
		} else {
			msg := handler.MakeNewMention(src) + " vamo a calmarno"
			telegrambot.SendMessage(bot, chatID, msg)
		}
	}
	// TODO: This is a quick fix and we should refactor it
	if !timekit.IsMidnight() {
		countPole = 0
		poleadoresIDs = nil
	}
}

func isPoleExhausted() bool {
	return countPole >= maxPoles
}

func haveINotSeenThisPoleadorID(poleadorID int64) bool {
	idStr := strconv.FormatInt(poleadorID, 10)
	if !stringkit.SliceContains(poleadoresIDs, idStr) {
		countPole++
		poleadoresIDs = append(poleadoresIDs, idStr)
		return true
	}
	return false
}

func handleMedal(src *tb.Message) string {
	switch countPole {
	case 1:
		return handler.MakeNewMention(src) + " ha hecho la pole. Medalla de oro! 🥇"
	case 2:
		return handler.MakeNewMention(src) + " ha hecho la subpole. Medalla de plata! 🥈"
	case 3:
		return handler.MakeNewMention(src) + " ha hecho un fail. Pa tu casa, champion"
	default:
		return "Error en la pole"
	}
}
