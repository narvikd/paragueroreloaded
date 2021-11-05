package envhandler

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"paraguero_reloaded/internal/logger"
	"paraguero_reloaded/internal/stringkit"
)

func LoadEnv(fileName string) {
	log := logger.GetLog()
	err := godotenv.Load(fileName)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetToken() (string, error) {
	env := os.Getenv("TOKEN")
	if !stringkit.IsStrEmpty(env) {
		return env, nil
	}
	return "", errors.New("\"TOKEN\" env variable was not found on the .env file")
}
