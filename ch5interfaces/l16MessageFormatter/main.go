package main

import "fmt"

type plainText struct {
	message string
}

func (f plainText) format() string {
	return f.message
}

type bold struct {
	message string
}

func (f bold) format() string {
	return fmt.Sprintf("**%s**", f.message)
}

type code struct {
	message string
}

func (c code) format() string {
	return fmt.Sprintf("`%s`", c.message)
}

type formatter interface {
	format() string
}

// Don't Touch below this line

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}
