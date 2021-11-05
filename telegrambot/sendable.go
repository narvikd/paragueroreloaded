package telegrambot

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/logger"
)

func SendMessage(bot *tb.Bot, chatID tb.ChatID, message string) {
	// options: markdown doesn't seem to work, instead it uses html.
	_, err := bot.Send(chatID, message, "html")
	logger.LogIfError(err, "error")
}

//func sendMessageToAdmin(bot *tb.Bot, message string, logIt bool) {
//	// This is hardcoded for now, since I'm the only admin.
//	chatID := tb.ChatID(1099020633)
//	SendMessage(bot, chatID, message)
//	log.Warnln("Message: " + message + ". Sent to admins.")
//	if logIt {
//		log.Warnln("Message: " + message + ". Sent to admins.")
//	}
//}

func SendSticker(bot *tb.Bot, chatID tb.Recipient, stickerID string) {
	_, err := bot.Send(chatID, &tb.Sticker{
		File: tb.File{FileID: stickerID},
	})
	logger.LogIfError(err, "error")
}

func SendKeyboard(bot *tb.Bot, route string, msgKeyboard string, menu *tb.ReplyMarkup) {
	bot.Handle(route, func(src *tb.Message) {
		chatID := tb.ChatID(src.Chat.ID)
		_, err := bot.Send(chatID, msgKeyboard, menu)
		logger.LogIfError(err, "error")
	})
}
