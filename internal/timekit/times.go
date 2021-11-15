package timekit

func IsDebugTime() bool {
	start := "17:50:00"
	end := "17:51:00"
	return isBetweenTimes(start, end)
}

func IsMidnight() bool {
	start := "00:00:00"
	end := "01:00:59"
	return isBetweenTimes(start, end)
}

func IsDawn() bool {
	start := "01:01:00"
	end := "05:59:59"
	return isBetweenTimes(start, end)
}

func IsMorning() bool {
	start := "06:00:00"
	end := "11:59:59"
	return isBetweenTimes(start, end)
}

func IsAfternoon() bool {
	start := "12:00:00"
	end := "17:59:59"
	return isBetweenTimes(start, end)
}

func IsEvening() bool {
	start := "18:00:00"
	end := "23:59:59"
	return isBetweenTimes(start, end)
}
