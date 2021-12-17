package telegrambot

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"paraguero_reloaded/internal/checkerr"
	"time"
)

// CreateBot creates a new bot object with the given token
func CreateBot(token string) *tb.Bot {
	// The 10 seconds timeout is the default and recommended setting, a lower number seems to affect performance.
	const pollerTimeout = 10

	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: pollerTimeout * time.Second},
	})
	checkerr.Fatalln(err)

	log.Println("Bot connected to Telegram Servers")
	return bot
}
