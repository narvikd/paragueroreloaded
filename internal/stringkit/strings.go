package stringkit

import "math/rand"

func RandomString(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func SliceContains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func IsStrEmpty(str string) bool {
	return str == ""
}
