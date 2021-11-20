package envhandler

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
	"paraguero_reloaded/internal/clearconsole"
	"paraguero_reloaded/internal/handleInterrupts"
	"paraguero_reloaded/internal/logger"
	"paraguero_reloaded/internal/osutils"
	"paraguero_reloaded/internal/stringkit"
	"path/filepath"
	"runtime"
	"strings"
)

var log = logger.GetLog()

// LoadEnv loads an .env file specified as a parameter
func LoadEnv(fileName string) {
	err := godotenv.Load(fileName)
	if err != nil {
		log.Errorln("Error loading .env file")
		token := readStringFromConsole("Enter token: ")
		if !isConsoleInputOK(token) {
			log.Errorln("Console input was not ok, retrying...")
			LoadEnv(fileName)
			return
		}
		writeTokenToFile(fileName, token)

		if wannaClearConsole() {
			clearconsole.CallClear()
		} else {
			fmt.Print("\n\n")
		}

		LoadEnv(fileName)
	}
}

// GetToken returns the token from the .env file. It should be used after LoadEnv is successful
func GetToken() (string, error) {
	env := os.Getenv("TOKEN")
	if !stringkit.IsStrEmpty(env) {
		return env, nil
	}
	return "", errors.New("\"TOKEN\" env variable was not found on the .env file")
}

func readStringFromConsole(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Errorln(errors.Wrap(err, "couldn't read console input"))
		if runtime.GOOS == "windows" && err.Error() == "EOF" {
			handleInterrupts.WindowsInterrupt <- true
		}
		readStringFromConsole(message)
	}

	// remove the delimiter from the string
	input = strings.TrimSuffix(input, "\n")
	if runtime.GOOS == "windows" {
		input = strings.TrimSuffix(input, "\r")
	}

	return input
}

func isConsoleInputOK(input string) bool {
	const tokenLength = 46
	return len(input) == tokenLength
}

func writeTokenToFile(fileName string, content string) {
	token := "TOKEN=" + content

	path, errPath := osutils.GetExecAbsPathFile(fileName)
	if errPath != nil {
		log.Fatalln(errPath)
	}

	// If the file doesn't exist, create it, or append to the file
	file, err := os.OpenFile(filepath.Clean(path), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "couldn't open token file"))
	}

	_, err = file.WriteString(token)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "couldn't write token string to file"))
	}

	err = file.Close()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "couldn't close the token file"))
	}
}

func wannaClearConsole() bool {
	inputConsole := readStringFromConsole("Do you want to clear the console? (y/n): ")
	input := strings.ToLower(inputConsole)
	switch input {
	case "y":
		return true
	case "n":
		return false
	default:
		log.Errorln("You must introduce \"y\" or \"n\"")
		wannaClearConsole()
	}
	return false
}
