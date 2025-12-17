package main

import (
	"fmt"
)

func main() {
	var x int = 50
	var y *int = &x
	fmt.Printf("ValueX:%v ValuePX:%v ValueY:%v ValuePY:%v ValueAY:%v\n", x, &x, y, &y, *y)
	*y = 100
	fmt.Printf("ValueX:%v ValuePX:%v ValueY:%v ValuePY:%v ValueAY:%v", x, &x, y, &y, *y)

}
