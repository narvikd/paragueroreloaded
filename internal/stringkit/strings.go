package stringkit

import (
	"paraguero_reloaded/internal/random"
)

// RandomString returns a random string from a slice
func RandomString(slice []string) string {
	rndNum := random.Num(len(slice))
	return slice[rndNum]
}

// SliceContains returns true if a string is in a slice
func SliceContains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

// IsStrEmpty returns true if a string is empty
func IsStrEmpty(str string) bool {
	return str == ""
}

func SplitStringNumTimes(str string, num int) []string {
	var words []string
	start := 0
	for i := 0; i < len(str); i++ {
		if i > 0 && i%num == 0 {
			if i < len(str) {
				words = append(words, str[start:i])
				start = i
			}
		}
	}
	if start < len(str) {
		words = append(words, str[start:])
	}
	return words
}
