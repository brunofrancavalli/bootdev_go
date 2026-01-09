package main

import (
	"fmt"
	"sync"
	"time"
)

func sendEmail(message string, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// Don't touch below this line

func test(message string, wg *sync.WaitGroup) {
	sendEmail(message, wg)
	// time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	test("Hello there Kaladin!", &wg)
	test("Hi there Shallan!", &wg)
	test("Hey there Dalinar!", &wg)

	wg.Wait()
}
