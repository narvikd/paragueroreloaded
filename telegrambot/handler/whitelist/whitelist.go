package whitelist

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

type TbFunc func(bot *tb.Bot, src *tb.Message)

func Only(f TbFunc) TbFunc {
	var ids = []int64{-1001155539490, -1001601989379, 1099020633}
	return func(bot *tb.Bot, src *tb.Message) {
		for _, id := range ids {
			if src.Chat.ID == id {
				f(bot, src)
			}
		}
	}
}
