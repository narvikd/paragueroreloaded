package kotlin

import (
    tb "gopkg.in/tucnak/telebot.v2"
    "paraguero_reloaded/internal/stringkit"
    "paraguero_reloaded/telegrambot"
    "strings"
)

var phrases = []string{
    "¿Que lenguaje se usará para programar plugins para IntelliJ?",
    "El único lenguaje moderno para hacer aplicaciones móviles",
    "Crea tu servidor http en solo 9 líneas con " +
        "<a href=\"https://ktor.io/\">Ktor</a>",
    "¿COMO? ¿qué no sabes como cambiar de lenguaje desde un proyecto gigante de Java 8? " +
        "pues claro, no tienes que hacerlo, usa Kotlin con su interoperabilidad",
    "Kotlin es el único lenguaje que hará que picar Spring no de asco",
    "¿Cansado de la documentación de Java? pues normal, pero para eso está Kotlin, " +
            "con su increible <a href=\"https://kotlinlang.org/docs/getting-started.html\">documentación</a>",
    "¿Sabías que no te hace falta escribir TypeScript? para algo existe " +
            "<a href=\"https://kotlinlang.org/docs/js-get-started.html#create-an-application\">KotlinJS</a>",
     "¿Te gusta React? no hay problema! "+
            "<a href=\"https://kotlinlang.org/docs/js-get-started.html#create-an-application\">KotlinJS for React</a>",
     "¿Quieres hacer binarios? no hay problema! "+
             "<a href=\"https://kotlinlang.org/docs/native-overview.html\">Kotlin Native</a>",
     "¿Echas de menos X libería de Java? pues importala!",
}

func GetKotlinMotivation(bot *tb.Bot, src *tb.Message) {
    cleanedReceivedMessage := strings.ToLower(src.Text)
    if cleanedReceivedMessage == "/kotlin" {
        chatID := tb.ChatID(src.Chat.ID)
        telegrambot.SendMessage(bot, chatID, stringkit.RandomString(phrases))
    }
}
