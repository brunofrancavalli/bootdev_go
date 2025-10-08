package main

import (
	"fmt"
)

func maxMessages(thresh int) int {
	total := 0
	for count := 0; ; count++ {
		total += 100 + count
		print(fmt.Sprintln(total))
	}

	return total
}
