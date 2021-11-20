package telegrambot

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/logger"
	"time"
)

var log = logger.GetLog()

// CreateBot creates a new bot object with the given token
func CreateBot(token string) *tb.Bot {
	// The 10 seconds timeout is the default and recommended setting, a lower number seems to affect performance.
	const pollerTimeout = 10
	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: pollerTimeout * time.Second},
	})
	logger.LogIfError(err, "fatal")
	log.Infoln("Bot connected to Telegram Servers")
	return bot
}

//func LogEndpointUsage(src *tb.Message, route string) {
//	logString := src.Sender.Username + "(" + strconv.Itoa(src.Sender.ID) + "): Used -> " + route + " <-."
//	log.Infoln(logString)
//}
