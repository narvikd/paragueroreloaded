package timekit

import (
	"fmt"
	"paraguero_reloaded/internal/logger"
	"time"
)

var log = logger.GetLog()

func isBetweenTimes(start, end string) bool {
	now, errNow := time.Parse("15:04:05", NowToHM()) // Get current time
	if errNow != nil {
		log.Errorln("Couldn't get current time" + errNow.Error())
		return false
	}

	startTime, errStartTime := time.Parse("15:04:05", start) // Get start time
	if errStartTime != nil {
		log.Errorln("Couldn't get start time" + errNow.Error())
		return false
	}

	endTime, errEndTime := time.Parse("15:04:05", end) // Get end time
	if errEndTime != nil {
		log.Errorln("Couldn't get end time" + errNow.Error())
		return false
	}

	// Check if current time is between start and end time
	if startTime.Before(endTime) {
		if now.After(startTime) && now.Before(endTime) {
			return true
		}
		return false
	}

	// If right now we're at a point in time after the endTime or before the startTime, return false
	if now.After(endTime) || now.Before(startTime) {
		return false
	}

	return true
}

// NowToHM returns the current time as a string with: "HH:MM:SS" as a format
func NowToHM() string {
	now := time.Now() // Get current time

	// Convert time to hours and minutes
	hours := now.Hour()
	minutes := now.Minute()
	seconds := now.Second()

	// Convert time to string
	hoursStr := timeToStr(hours)
	minutesStr := timeToStr(minutes)
	secondsStr := timeToStr(seconds)

	// Return it
	hoursMinutesStr := hoursStr + ":" + minutesStr + ":" + secondsStr
	return hoursMinutesStr
}

func timeToStr(time int) string {
	if time < 10 {
		return "0" + fmt.Sprintf("%d", time)
	}
	return fmt.Sprintf("%d", time)
}
