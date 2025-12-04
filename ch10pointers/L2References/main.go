package main

import (
	"strings"
)

func removeProfanity(message *string) {
	newMessage := message
	words := []string{"fubb", "shiz", "witch"}

	for _, word := range words {

		*newMessage = strings.Replace(*newMessage, word, getAsterisks(word), -1)
	}
}

func getAsterisks(value string) string {
	asterisks := ""

	for index := 0; index < len(value); index++ {
		asterisks += "*"
	}

	return asterisks
}
