package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var i = 0.0
	var reverseInt = 0
	for ok := int(x); ok != 0; ok = x / 10 {
		reverseInt += int(int(ok%10) * int(math.Pow(10, i)))
		i++
	}
	return reverseInt
}

func main() {
	fmt.Println("start")
	defer fmt.Println("finish")

	fmt.Println(reverse(216))
}
