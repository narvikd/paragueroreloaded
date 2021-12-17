package main

import (
	"paraguero_reloaded/internal/checkerr"
	"paraguero_reloaded/internal/config/configfile"
	"paraguero_reloaded/telegrambot"
	"paraguero_reloaded/telegrambot/endpoint"
)

func main() {
	token, err := configfile.GetToken("env.env")
	checkerr.Fatalln(err)

	bot := telegrambot.CreateBot(token)
	endpoint.LoadEndpoints(bot)
	bot.Start()
}
