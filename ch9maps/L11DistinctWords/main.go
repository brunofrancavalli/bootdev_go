package main

import "strings"

func countDistinctWords(messages []string) int {

	counterMap := make(map[string]int)
	for _, message := range messages {
		words := strings.Split(message, " ")
		for _, word := range words {
			wordLower := strings.ToLower(word)
			if len(wordLower) > 0 {
				_, ok := counterMap[wordLower]
				if !ok {
					counterMap[wordLower] = 1
				} else {
					counterMap[wordLower] = 1 + counterMap[wordLower]
				}
			}
		}
	}

	return len(counterMap)
}
