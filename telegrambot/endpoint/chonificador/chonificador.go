package chonificador

import (
	"github.com/dlclark/regexp2"
	"github.com/valyala/fasthttp"
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot"
	"strconv"
	"strings"
	"time"
)

func TranslateToChoni(bot *tb.Bot, src *tb.Message) {
	if strings.HasPrefix(src.Text, "/choni ") {
		chatID := tb.ChatID(src.Chat.ID)
		toTranslate := strings.SplitAfter(src.Text, "/choni ")

		client := fasthttp.Client{} // Create fasthttp client
		// Get request/response from pool
		req := fasthttp.AcquireRequest()
		res := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(res)

		apiEndpoint := "https://gomorria.webcindario.com/chonificador.php"
		req.SetRequestURI(apiEndpoint)
		// Sets user agent and content type (don't change it) in order for POST to work
		req.Header.Add("User-Agent", "Paraguas Bot")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.SetMethod("POST")
		// Sets the body message as in key1=value&key=value2
		req.SetBodyString("t=" + toTranslate[1] + "&mayus=1")

		// Make request
		if err := client.DoTimeout(req, res, 5*time.Second); err != nil {
			telegrambot.SendMessage(bot, chatID, "Ayy mi arma esa riqüest es muy lenta mi arma")
		}

		// Check statuscode
		if res.StatusCode() != 200 {
			status := strconv.Itoa(res.StatusCode())
			telegrambot.SendMessage(bot, chatID, "Mi arma ha ocurriu un errur con la riqüest. Errur: "+status)
		}

		// Formats html with regex because >tfw no good soup library
		regex := regexp2.MustCompile("(?<=<textarea name=\"t\" id=\"resultadoa\">)(.*)(<\\/textarea>)", 0)
		match, _ := regex.FindStringMatch(string(res.Body()))

		// The output we want it's in group 1
		group := match.Groups()
		telegrambot.SendMessage(bot, chatID, group[1].String())
	}
}
