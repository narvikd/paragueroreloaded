package debugEndpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/logger"
	"paraguero_reloaded/telegrambot"
	"paraguero_reloaded/telegrambot/handler/ontext"
	"strconv"
)

var log = logger.GetLog()

func GetStickerInfo(bot *tb.Bot, src *tb.Message) {
	chatID := tb.ChatID(src.Chat.ID)
	telegrambot.SendMessage(bot, chatID, src.Sticker.UniqueID)
	log.Infoln(src.Sticker.UniqueID)
}

func GetSenderID(bot *tb.Bot) {
	ontext.AddEndpoint(bot, func(bot *tb.Bot, src *tb.Message) {
		chatID := tb.ChatID(src.Chat.ID)
		if src.IsForwarded() {
			str := strconv.Itoa(src.OriginalSender.ID)
			telegrambot.SendMessage(bot, chatID, str)
		} else {
			str := strconv.Itoa(src.Sender.ID)
			telegrambot.SendMessage(bot, chatID, str)
		}

		log.Warnln(src.OriginalSender.ID)
	})
}

func GetUsername(bot *tb.Bot) {
	ontext.AddEndpoint(bot, func(bot *tb.Bot, src *tb.Message) {
		log.Warnln(src.Sender.Username)
	})
}
