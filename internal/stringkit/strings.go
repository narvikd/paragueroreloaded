package stringkit

import "math/rand"

// RandomString returns a random string from a slice
func RandomString(slice []string) string {
	return slice[rand.Intn(len(slice))]
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
