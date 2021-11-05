package envhandler

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"paraguero_reloaded/internal/logger"
)

func LoadEnv(fileName string) {
	log := logger.GetLog()
	err := godotenv.Load(fileName)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetToken(environmentVar string) (string, error) {
	switch environmentVar {
	case "PROD":
		return os.Getenv("TOKEN_PROD"), nil
	case "DEV":
		return os.Getenv("TOKEN_DEV"), nil
	default:
		return "", errors.New("environment not set correctly")
	}
}
