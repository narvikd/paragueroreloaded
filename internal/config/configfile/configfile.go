package configfile

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
)

// GetToken returns the token from the .env file.
func GetToken(fileName string) (string, error) {
	err := godotenv.Load(fileName)
	if err != nil {
		return "", errors.Wrap(err, "error loading config file")
	}

	env := os.Getenv("TOKEN")
	if env == "" {
		return "", errors.New("\"TOKEN\" env variable was not found on the .env file")
	}
	return env, nil
}
