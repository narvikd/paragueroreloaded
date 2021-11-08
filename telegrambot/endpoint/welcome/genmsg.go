package welcomeendpoint

import (
	"paraguero_reloaded/internal/stringkit"
)

var phrases = []string{
	"del Imperio que crea software",
	"de Jesucristo programador",
	"de Java",
	"de Python",
	"de Román el europeo",
	"de Moldavia",
	"de la ingeniera rusa",
	"de Cecilio",
	"de Ayuso",
	"de Leetcode",
	"del Clean Code",
	"de las apps basura del Play Store",
	"de España",
	"del Espíritu Santo",
	"de la España programadora",
}

func getWelcomeMsg(senderMention string) string {
	newPhrases := genWelcomeMsg()
	message := "Por la gloria " +
		newPhrases[0] + ", " +
		newPhrases[1] + " y " +
		newPhrases[2] + ", " +
		"yo te bendigo y te doy la bienvenida, " +
		senderMention
	return message
}

func genWelcomeMsg() []string {
	var msgSlice []string
	for len(msgSlice) < 3 {
		rndStr := stringkit.RandomString(phrases)
		if !stringkit.SliceContains(msgSlice, rndStr) {
			msgSlice = append(msgSlice, rndStr)
		}
	}
	return msgSlice
}
