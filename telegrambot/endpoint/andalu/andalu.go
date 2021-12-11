package andalu

import (
	"encoding/json"
	"errors"
	"git.dani.icu/narvikd/tempfasthttp"
	tb "gopkg.in/tucnak/telebot.v2"
	"net/http"
	url2 "net/url"
	"paraguero_reloaded/telegrambot"
	"strconv"
	"strings"
	"time"
)

// TranslateToAndalu translates ugly Spanish syntax into the language of the Gods
func TranslateToAndalu(bot *tb.Bot, src *tb.Message) {
	if strings.HasPrefix(src.Text, "/andalu ") {
		toTranslate := strings.SplitAfter(src.Text, "/andalu ")
		chatID, model, _ := Common(bot, src, toTranslate[1])
		telegrambot.SendMessage(bot, chatID, model.Andaluh)
	}
}

func Common(bot *tb.Bot, src *tb.Message, toTranslate string) (tb.ChatID, *Andalu, error) {
	chatID := tb.ChatID(src.Chat.ID)

	client := fasthttp.Client{} // Create fasthttp client
	// Get request/response from pool
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	apiEndpoint := "https://api.andaluh.es/epa?spanish="
	url := apiEndpoint + url2.QueryEscape(toTranslate)
	req.SetRequestURI(url)

	// Make request
	const requestTimeout = 5
	if err := client.DoTimeout(req, res, requestTimeout*time.Second); err != nil {
		telegrambot.SendMessage(bot, chatID, "Lo siento mi arma etoy de zieta, luego me dise ok")
		return 0, nil, err
	}

	// Check statuscode
	if res.StatusCode() != http.StatusOK {
		status := strconv.Itoa(res.StatusCode())
		telegrambot.SendMessage(bot, chatID, "Mi arma ha ocurriu un errur con la riqüest. Errur: "+status)
		return 0, nil, errors.New("status code not 200")
	}

	// Unmarshal the andalu
	model := new(Andalu)
	err := json.Unmarshal(res.Body(), &model)
	if err != nil {
		telegrambot.SendMessage(bot, chatID, "Killu nu puedo parzea esto")
		return 0, nil, err
	}

	// If everything was ok, return the model
	return chatID, model, nil
}
