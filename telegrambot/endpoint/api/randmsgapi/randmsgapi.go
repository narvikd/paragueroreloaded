package randmsgapi

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot"
	"paraguero_reloaded/telegrambot/handler/logger"
	"strconv"
	"time"
)

type ApiModel struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func Get(bot *tb.Bot, route string) {
	bot.Handle(route, func(src *tb.Message) {
		chatID := tb.ChatID(src.Chat.ID)
		msg, _ := Dial(route)
		if msg != "" {
			logger.Logger(src)
			telegrambot.SendMessage(bot, chatID, msg)
		}
	})
}

func Dial(route string) (string, error) {
	client := fasthttp.Client{}
	// Get request/response from pool
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	url := "http://api:3001/msg" + route
	req.SetRequestURI(url)

	// Make request
	timeout := 3 * time.Second
	if err := client.DoTimeout(req, res, timeout); err != nil {
		return "", errors.Wrap(err, "Dial connection failed, timeout reached. Api could be down")
	}

	// Check status code
	if res.StatusCode() != 200 {
		return "", errors.New("api response status code " + strconv.Itoa(res.StatusCode()))
	}

	model := new(ApiModel)
	err := json.Unmarshal(res.Body(), &model)
	if err != nil {
		return "", errors.Wrap(err, "couldn't parse api json")
	}

	if !model.Success {
		errMsg := "api no success. api response status code: " + strconv.Itoa(res.StatusCode()) + ". message: " + model.Message
		return "", errors.New(errMsg)
	}

	return model.Data, nil
}
