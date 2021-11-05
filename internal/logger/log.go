package logger

import (
	"github.com/sirupsen/logrus"
	"strings"
)

func GetLog() *logrus.Logger {
	var log = logrus.New()
	// This format is not to be changed to your current time, this is according to the constants at: https://golang.org/pkg/time/#pkg-constants
	log.SetFormatter(&logrus.TextFormatter{TimestampFormat: "02-01-2006 15:04:05 MST", FullTimestamp: true})
	return log
}

func LogIfError(err error, errType string) {
	log := GetLog()
	if err != nil {
		switch strings.ToLower(errType) {
		case "fatal":
			log.Fatalln(err)
		case "panic":
			log.Panicln(err)
		case "error":
			log.Errorln(err)
		case "warn":
			log.Warnln(err)
		case "debug":
			log.Debugln(err)
		case "info":
			log.Infoln(err)
		default:
			log.Errorln("Error: " + err.Error() + ". Was logged, but the type was unknown.")
		}
	}
}
