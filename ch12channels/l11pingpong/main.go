package main

import (
	"fmt"
)

func pingPong(numPings int) {
	done := make(chan int)
	pings := make(chan int)
	pongs := make(chan int)
	go ponger(pings, pongs, done)
	go pinger(pings, pongs, numPings)

	<-done
}

// don't touch below this line

func pinger(pings chan int, pongs chan int, numPings int) {
	for i := 0; i < numPings; i++ {
		fmt.Printf("sending ping %v\n", i)
		pings <- i
		iReturn, _ := <-pongs

		if i != iReturn {
			fmt.Println("Did not get a pong back")
		} else {
			fmt.Printf("Got pong %v\n", iReturn)
		}

	}
	close(pings)
}

func ponger(pings, pongs, done chan int) {
	for {
		i, ok := <-pings
		if ok {
			fmt.Printf("got ping %v, sending pong %v\n", i, i)
			pongs <- i
		} else {
			break
		}
	}
	fmt.Println("pings done")
	close(pongs)
	close(done)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	test(4)
	test(3)
	test(2)
	fmt.Println("Application Done")
}
