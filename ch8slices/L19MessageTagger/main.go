package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for index, _ := range messages {
		messages[index].tags = make([]string, 0)
		messages[index].tags = append(messages[index].tags, tagger(messages[index])...)
		fmt.Println(len(messages[index].tags))
	}

	return messages
}

func tagger(msg sms) []string {
	tags := []string{}

	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}
