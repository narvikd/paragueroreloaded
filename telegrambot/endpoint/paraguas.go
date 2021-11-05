package endpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/stringkit"
	"paraguero_reloaded/telegrambot"
	"strings"
)

// paraguas is an endpoint that executes if the "Abriendo paraguas" regex is triggered
func paraguas(bot *tb.Bot, src *tb.Message) {
	const regexPattern = "abr(e|o|iendo) paraguas"
	cleanedReceivedMessage := strings.ToLower(src.Text)

	if stringkit.IsRegexMatch(cleanedReceivedMessage, regexPattern) {
		chatID := tb.ChatID(src.Chat.ID)
		const stickerID = "CAACAgIAAxkBAAIBIWGDDf3G7veVZLqqbKxh5hPiMuTNAAJaAAMsyqoHJWwfhyyBODEhBA"

		telegrambot.SendMessage(bot, chatID, "Abriendo paraguas")
		telegrambot.SendSticker(bot, chatID, stickerID)
	}
}

// paraguasSticker is an endpoint that executes if a paraguas sticker is received
func paraguasSticker(bot *tb.Bot, src *tb.Message) {
	var stickers = []string{
		"AgADWgADLMqqBw",
		"AgADPwsAAsjB4Co",
		"AgADQQEAAqghIQY",
	}

	if stringkit.SliceContains(stickers, src.Sticker.UniqueID) {
		chatID := tb.ChatID(src.Chat.ID)
		telegrambot.SendMessage(bot, chatID, "Abriendo paraguas")
	}
}
