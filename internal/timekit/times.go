package timekit

func IsBetweenTimesBool(start string, end string) {
	returnFalseIfErr(IsBetweenTimes(start, end))
}

func IsMidnight() bool {
	start := "00:00:00"
	end := "01:00:59"
	return returnFalseIfErr(IsBetweenTimes(start, end))
}

func IsDawn() bool {
	start := "01:01:00"
	end := "05:59:59"
	return returnFalseIfErr(IsBetweenTimes(start, end))
}

func IsMorning() bool {
	start := "06:00:00"
	end := "11:59:59"
	return returnFalseIfErr(IsBetweenTimes(start, end))
}

func IsAfternoon() bool {
	start := "12:00:00"
	end := "17:59:59"
	return returnFalseIfErr(IsBetweenTimes(start, end))
}

func IsEvening() bool {
	start := "18:00:00"
	end := "23:59:59"
	return returnFalseIfErr(IsBetweenTimes(start, end))
}

// returnFalseIfErr is a function I had to make in case I'm programming at 2 AM and I can't distinguish between left and right
func returnFalseIfErr(err error) bool {
	if err != nil {
		return false
	}
	return true
}
