package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed {
		return 2 * len(e.body)
	} else {
		return 5 * len(e.body)
	}

}

func (e email) format() string {
	subscribedText := "Not Subscribed"
	if e.isSubscribed {
		subscribedText = "Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, subscribedText)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
