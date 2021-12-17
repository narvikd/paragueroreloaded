package configfile

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
)

// GetToken returns the token from the .env file. It should be used after loadEnv is successful
func GetToken(fileName string) (string, error) {
	err := loadEnv(fileName)
	if err != nil {
		return "", err
	}

	env := os.Getenv("TOKEN")
	if env != "" {
		return env, nil
	}
	return "", errors.New("\"TOKEN\" env variable was not found on the .env file")
}

// loadEnv loads an .env file specified as a parameter
func loadEnv(fileName string) error {
	err := godotenv.Load(fileName)
	return errors.Wrap(err, "error loading config file")
}
