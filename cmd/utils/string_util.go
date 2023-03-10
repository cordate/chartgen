package utils

import "strings"

func IsEmpty(str string) bool {
	str = strings.TrimSpace(str)
	return len(str) == 0
}

func NotEmpty(str string) bool {
	return !IsEmpty(str)
}

func ArrayToString(array []string) string {
	return strings.Join(array, ",")
}
