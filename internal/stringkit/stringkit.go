package stringkit

import "regexp"

// SliceContains returns true if a string is in a slice
func SliceContains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

// IsRegexMatch returns true if the string matches the regexp.
func IsRegexMatch(str string, pattern string) bool {
	if ok, _ := regexp.MatchString(pattern, str); ok {
		return true
	}
	return false
}
