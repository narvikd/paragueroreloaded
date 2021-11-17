package andalu

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	tb "gopkg.in/tucnak/telebot.v2"
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
		chatID, model := Common(bot, src, toTranslate[1])
		telegrambot.SendMessage(bot, chatID, model.Andaluh)
	}
}

func Common(bot *tb.Bot, src *tb.Message, toTranslate string) (tb.ChatID, *Andalu) {
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
	if err := client.DoTimeout(req, res, 5*time.Second); err != nil {
		telegrambot.SendMessage(bot, chatID, "Ayy mi arma esa riqüest es muy lenta mi arma")
	}

	// Check statuscode
	if res.StatusCode() != 200 {
		status := strconv.Itoa(res.StatusCode())
		telegrambot.SendMessage(bot, chatID, "Mi arma ha ocurriu un errur con la riqüest. Errur: "+status)
	}

	// Unmarshal the andalu
	model := new(Andalu)
	err := json.Unmarshal(res.Body(), &model)
	if err != nil {
		telegrambot.SendMessage(bot, chatID, "Killu nu puedo parzea esto")
	}

	// If everything was ok, return the model
	return chatID, model
}
