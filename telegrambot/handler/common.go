package handler

import tb "gopkg.in/tucnak/telebot.v2"

func IterateEndpoints(bot *tb.Bot, src *tb.Message, slice []func(bot *tb.Bot, src *tb.Message)) {
	for _, function := range slice {
		function(bot, src)
	}
}
