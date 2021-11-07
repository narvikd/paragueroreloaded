package debugEndpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/logger"
	"paraguero_reloaded/internal/timekit"
	"paraguero_reloaded/telegrambot"
	"paraguero_reloaded/telegrambot/handler"
	"paraguero_reloaded/telegrambot/handler/ontext"
	"strconv"
)

var log = logger.GetLog()
var msgCopy string
var counterCopyMsg int

func GetStickerInfo(bot *tb.Bot, src *tb.Message) {
	chatID := tb.ChatID(src.Chat.ID)
	telegrambot.SendMessage(bot, chatID, src.Sticker.UniqueID)
	log.Infoln(src.Sticker.UniqueID)
}

func GetSenderID(bot *tb.Bot, src *tb.Message) {
	chatID := tb.ChatID(src.Chat.ID)
	str := strconv.Itoa(src.Sender.ID)
	msg := handler.MakeNewMention(src) + " tu ID de telegram es: " + str
	if src.IsForwarded() {
		str = strconv.Itoa(src.OriginalSender.ID)
	}
	telegrambot.SendMessage(bot, chatID, msg)
}

func GetUsername(bot *tb.Bot) {
	ontext.AddEndpoint(bot, func(bot *tb.Bot, src *tb.Message) {
		log.Warnln(src.Sender.Username)
	})
}

// GetCurrentTime gets the current time expressed as HH:MM, only the admin user can launch it
func GetCurrentTime(bot *tb.Bot, route string) {
	bot.Handle(route, func(src *tb.Message) {
		if telegrambot.IsAdmin(src) {
			chatID := tb.ChatID(src.Chat.ID)
			telegrambot.SendMessage(bot, chatID, timekit.NowToHM())
		}
	})
}

func PrintSrc(bot *tb.Bot, route string) {
	bot.Handle(route, func(src *tb.Message) {
		chatID := tb.ChatID(src.Chat.ID)

		srcCopy := *src // https://stackoverflow.com/a/38443432
		log.Warn("src: ")
		log.Warnln(src)

		msgOriginal := "He procesado: " + strconv.Itoa(srcCopy.ID) + " mensajes desde mi creación"
		if counterCopyMsg == 0 {
			msgCopy = "He procesado: " + strconv.Itoa(src.ID) + " mensajes desde mi creación"
		}
		counterCopyMsg++

		telegrambot.SendMessage(bot, chatID, "(Original) "+msgOriginal)
		telegrambot.SendMessage(bot, chatID, "(Copy) "+msgCopy)
	})
}
