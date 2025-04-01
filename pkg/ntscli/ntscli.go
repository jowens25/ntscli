package ntscli

import (
	"strconv"
)

func Get(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
	}
	return result
}

func Set(input string) (kind string) {

	return "digit"
}

func inspectNumbers(input string) (count int) {
	for _, c := range input {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			count++
		}
	}
	return count
}
