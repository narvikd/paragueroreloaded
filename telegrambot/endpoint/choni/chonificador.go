package choni

import (
	"errors"
	"git.dani.icu/narvikd/tempfasthttp"
	"github.com/dlclark/regexp2"
	tb "gopkg.in/tucnak/telebot.v2"
	"net/http"
	"paraguero_reloaded/telegrambot"
	"strconv"
	"strings"
	"time"
)

func TranslateToChoni(bot *tb.Bot, src *tb.Message) {
	if strings.HasPrefix(src.Text, "/choni ") {
		toTranslate := strings.SplitAfter(src.Text, "/choni ")
		chatID, msg, _ := Common(bot, src, toTranslate[1])
		telegrambot.SendMessage(bot, chatID, msg)
	}
}

func Common(bot *tb.Bot, src *tb.Message, toTranslate string) (tb.ChatID, string, error) {
	chatID := tb.ChatID(src.Chat.ID)

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
	req.SetBodyString("t=" + toTranslate + "&mayus=1")

	// Make request
	const requestTimeout = 5
	if err := client.DoTimeout(req, res, requestTimeout*time.Second); err != nil {
		telegrambot.SendMessage(bot, chatID, "loOh çieNtoooh miIh aRMaaH etoy dEeH çiEtaaaH, lueGooh meeH Diçeeh OK")
		return 0, "", err
	}

	// Check statuscode
	if res.StatusCode() != http.StatusOK {
		status := strconv.Itoa(res.StatusCode())
		telegrambot.SendMessage(bot, chatID, "miiIH arMaaah hAah OkurriuuUh un ErruR kon laAh riqüest. Errur: "+status)
		return 0, "nil", errors.New("status code not 200")
	}

	// Formats html with regex because >tfw no good soup library
	regex := regexp2.MustCompile("(?<=<textarea name=\"t\" id=\"resultadoa\">)(.*)(<\\/textarea>)", 0)
	match, _ := regex.FindStringMatch(string(res.Body()))

	// The output we want it's in group 1
	group := match.Groups()

	message := group[1].String()
	return chatID, message, nil
}
