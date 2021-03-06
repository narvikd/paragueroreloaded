package logger

import (
	"github.com/TwiN/go-color"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"paraguero_reloaded/telegrambot/handler"
	"strconv"
)

func Register(bot *tb.Bot, src *tb.Message) {
	Logger(src)
}

func Logger(src *tb.Message) {
	chat := "\"" + color.Ize(color.Red, src.Chat.Title) + "\": "
	user := color.Ize(color.Green, handler.MakeNewMention(src)) + ". "
	used := "said \"" + color.Ize(color.Bold+color.Red, src.Text) + "\". "
	chatID := "(ChatID: \"" + color.Ize(color.Purple, strconv.FormatInt(src.Chat.ID, 10)) + "\". "
	chatType := "Chat type: " + color.Ize(color.Purple, string(src.Chat.Type)) + ") "
	str := chat + user + used + chatID + chatType
	log.Println(str)
}
