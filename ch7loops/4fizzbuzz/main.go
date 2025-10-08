package main

import (
	"fmt"
)

func fizzbuzz() {
	for count := 0; count <= 100; count++ {
		remainder3 := count % 3
		remainder5 := count % 5

		if remainder3 == 0 || remainder5 == 0 {
			output := ""
			if remainder3 == 0 {
				output = "fizz"
			}
			if remainder5 == 0 {
				output += "buzz"
			}
			fmt.Println(output)
		} else {
			fmt.Println(count)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
