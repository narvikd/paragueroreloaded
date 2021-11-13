package rust

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/internal/stringkit"
	"paraguero_reloaded/telegrambot"
	"strings"
)

var phrases = []string{
	"Puedes crear maravillosas interfaces GUI con " +
		"<a href=\"https://github.com/linebender/druid\">Druid</a>!",
	"Puedes crear maravillosas interfaces TUI con " +
		"<a href=\"https://github.com/clap-rs/clap\">Clap</a>!",
	"Puedes crear maravillosos juegos/gráficos con <b>Game Engines</b> como: " +
		"<a href=\"https://amethyst.rs/\">Amethyst</a>, " +
		"<a href=\"https://bevyengine.org/\">Bevy</a> y " +
		"<a href=\"https://ggez.rs/\">Ggez</a> (entre muchos otros).",
	"Lenguaje piola.",
	"Este señor sabe.",
	"Es el lenguaje más amado, temido y desdeado según la " +
		"<a href=\"https://insights.stackoverflow.com/survey/2016#technology-most-loved-dreaded-and-wanted\">" +
		"<b>Stackoverflow Dev Survey</b></a>",
	"¿Te interesa programar X? Claro que puedes hacerlo con Rust :D.",
	"Ferris, la mascota más bonita jamás creada.",
	"Null, ¿Qué es null? ¿Suena a invento de mortales?",
	"Al no tener GC, sus tiempos de ejecución van to follaos.",
	"El camino puede ser duro, pero el único lifetime que importa es el de tu constancia.",
	"Es un lenguaje de dioses, por eso mi amo es un mierdas.",
}

func GetRustMotivation(bot *tb.Bot, src *tb.Message) {
	cleanedReceivedMessage := strings.ToLower(src.Text)
	if cleanedReceivedMessage == "/rust" {
		chatID := tb.ChatID(src.Chat.ID)
		telegrambot.SendMessage(bot, chatID, stringkit.RandomString(phrases))
	}
}
