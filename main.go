package main

import (
	"paraguero_reloaded/internal/appinit"
	"paraguero_reloaded/internal/envhandler"
	"paraguero_reloaded/internal/logger"
	"paraguero_reloaded/telegrambot"
	"paraguero_reloaded/telegrambot/endpoint"
)

const environment = "PROD"

func init() {
	appinit.SeedRNG()
	envhandler.LoadEnv("env.env")
}

func main() {
	token, err := envhandler.GetToken(environment)
	logger.LogIfError(err, "fatal")
	bot := telegrambot.CreateBot(token)
	endpoint.LoadEndpoints(bot)
	bot.Start()
}
