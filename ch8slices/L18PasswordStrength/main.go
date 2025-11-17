package main

import (
	"strings"
)

func isValidPassword(password string) bool {
	isValid := true

	passwordRune := []rune(password)
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	if len(password) < 5 || len(password) > 12 {
		isValid = false
	} else {
		if strings.ToLower(password) == password {
			isValid = false
		}
	}

	oneDigit := false

	for _, character := range passwordRune {

		for _, digit := range digits {
			if digit == character {
				oneDigit = true
			}
		}
	}

	return isValid && oneDigit
}
