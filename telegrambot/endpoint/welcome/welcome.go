package welcome

import (
	"github.com/pkg/errors"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"paraguero_reloaded/internal/checkerr"
	"paraguero_reloaded/internal/stringkit"
	"paraguero_reloaded/telegrambot/endpoint/api/randmsgapi"
	"paraguero_reloaded/telegrambot/handler"
)

func Run(bot *tb.Bot, src *tb.Message) {
	chatID := tb.ChatID(src.Chat.ID)
	senderMention := handler.MakeNewMention(src)
	_, err := bot.Send(chatID, getWelcomeMsg(senderMention), "html")
	checkerr.Println(errors.Wrap(err, "couldn't send Run"))
}

func getWelcomeMsg(senderMention string) string {
	newPhrases, err := genWelcomeMsg()
	if err != nil {
		log.Println(err)
		return "Bienvenido " + senderMention
	}
	message := "Por la gloria " +
		newPhrases[0] + ", " +
		newPhrases[1] + " y " +
		newPhrases[2] + ", " +
		"yo te bendigo y te doy la bienvenida, " +
		senderMention
	return message
}

func genWelcomeMsg() ([]string, error) {
	var msgSlice []string
	for len(msgSlice) < 3 {
		rndStr, err := randmsgapi.Dial("/welcome")
		if err != nil {
			return nil, errors.New("couldn't generate new welcome msg")
		}
		if !stringkit.SliceContains(msgSlice, rndStr) {
			msgSlice = append(msgSlice, rndStr)
		}
	}
	return msgSlice, nil
}
