package main

import (
	"fmt"
)

func printPrimes(max int) {
	for count := 2; count < max; count++ {
		if count == 2 {
			fmt.Println(count)
		} else if count%2 == 0 {
			continue
		} else {
			broke := false
			for count2 := 2; count2 < count; count2++ {
				if (count % count2) == 0 {
					broke = true
					break
				}
			}

			if !broke {
				fmt.Println(count)
			}
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
