package stringkit

import "regexp"

// IsRegexMatch exists because
// "if matchString, _ := regexp.MatchString(pattern, cleanedReceivedMessage); matchString"
// is ugly as sin
func IsRegexMatch(str string, pattern string) bool {
	if ok, _ := regexp.MatchString(pattern, str); ok {
		return true
	}
	return false
}
