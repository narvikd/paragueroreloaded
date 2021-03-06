package endpoint

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot"
)

// links is an endpoint for the (default) /links command.
func links(bot *tb.Bot, route string) {
	menu := bot.NewMarkup()
	menu.Inline(
		menu.Row(
			menu.URL("Hilo Forocoches 🚗", "https://www.forocoches.com/foro/showthread.php?t=8055773"),
			menu.URL("¿Quien soy? 🐸", "https://github.com/narvikd/paragueroreloaded"),
		),
		menu.Row(
			menu.URL("Enlace grupo 🔗", "https://bit.ly/dawdamfc"),
			menu.URL("Offtopic (+18) 😈", "https://t.me/joinchat/kZfRLk6IXO0wNWE0"),
		),
	)
	telegrambot.SendKeyboard(bot, route, "Aquí los tienes, vago de mierda.", menu)
}
