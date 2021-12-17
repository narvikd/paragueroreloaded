package telegrambot

import (
	"github.com/pkg/errors"
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/checkerr"
)

// SendMessage sends a message to the chat, both have to be specified as a parameter
func SendMessage(bot *tb.Bot, chatID tb.ChatID, message string) {
	// options: markdown doesn't seem to work, instead it uses html.
	_, err := bot.Send(chatID, message, "html")
	checkerr.Println(errors.Wrap(err, "couldn't send message"))
}

// SendSticker sends a sticker to the chat, both have to be specified as a parameter
func SendSticker(bot *tb.Bot, chatID tb.Recipient, stickerID string) {
	_, err := bot.Send(chatID, &tb.Sticker{
		File: tb.File{FileID: stickerID},
	})
	checkerr.Println(errors.Wrap(err, "couldn't send sticker"))
}

// SendKeyboard sends a keyboard to the chat.
// The route and the message that are going to be sent before the keyboard have to be specified as a parameter.
// the menu is a bot.NewMarkup().Inline object
func SendKeyboard(bot *tb.Bot, route string, msgKeyboard string, menu *tb.ReplyMarkup) {
	bot.Handle(route, func(src *tb.Message) {
		chatID := tb.ChatID(src.Chat.ID)
		_, err := bot.Send(chatID, msgKeyboard, menu)
		checkerr.Println(errors.Wrap(err, "couldn't send keyboard"))
	})
}
