package main

import "strings"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for msgIndex, msgItem := range msg {
		for _, badWord := range badWords {
			if strings.Contains(strings.ToLower(badWord), strings.ToLower(msgItem)) {
				return msgIndex
			}
		}
	}

	return -1
}
