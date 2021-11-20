package main

import (
	"paraguero_reloaded/internal/appinit"
	"paraguero_reloaded/internal/envhandler"
	"paraguero_reloaded/internal/handleInterrupts"
	"paraguero_reloaded/internal/logger"
	"paraguero_reloaded/telegrambot"
	"paraguero_reloaded/telegrambot/endpoint"
)

func init() {
	appinit.SeedRNG()
	go handleInterrupts.HandleStop()
	envhandler.LoadEnv("env.env")
}

func main() {
	token, err := envhandler.GetToken()
	logger.LogIfError(err, "fatal")
	bot := telegrambot.CreateBot(token)
	endpoint.LoadEndpoints(bot)
	bot.Start()
}
