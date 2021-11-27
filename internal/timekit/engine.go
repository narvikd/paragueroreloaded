package timekit

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

func IsBetweenTimes(start, end string) error {
	now, errNow := time.Parse("15:04:05", NowToHM()) // Get current time
	if errNow != nil {
		return errors.Wrap(errNow, "couldn't get current time")
	}

	startTime, errStartTime := time.Parse("15:04:05", start) // Get start time
	if errStartTime != nil {
		return errors.Wrap(errStartTime, "couldn't get start time")
	}

	endTime, errEndTime := time.Parse("15:04:05", end) // Get end time
	if errEndTime != nil {
		return errors.Wrap(errEndTime, "couldn't get end time")
	}

	// Check if current time is between start and end time
	if startTime.Before(endTime) {
		if now.After(startTime) && now.Before(endTime) {
			return nil
		}
		return errors.New("current time isn't between start and end time")
	}

	// If right now we're at a point in time after the endTime or before the startTime
	if now.After(endTime) || now.Before(startTime) {
		return errors.New("current time after endTime or before startTime")
	}

	return nil
}

// NowToHM returns the current time as a string with: "HH:MM:SS" as a format
func NowToHM() string {
	now := time.Now() // Get current time

	// Convert time to hours and minutes
	hours := now.Hour()
	minutes := now.Minute()
	seconds := now.Second()

	// Convert time to string
	hoursStr := TimeToStr(hours)
	minutesStr := TimeToStr(minutes)
	secondsStr := TimeToStr(seconds)

	// Return it
	hoursMinutesStr := hoursStr + ":" + minutesStr + ":" + secondsStr
	return hoursMinutesStr
}

/*
	TimeToStr converts an int (time.Hour for example) into a string.

	If we have 1 hours and 1 seconds, we don't want to make a date as: "1:1". But 01:01.

	That's why it checks for digits lower than 10.

	Warning: Don't use it with non-validated times like "61 seconds".
*/
func TimeToStr(time int) string {
	if time < 10 {
		return "0" + fmt.Sprintf("%d", time)
	}
	return fmt.Sprintf("%d", time)
}
