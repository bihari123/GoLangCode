package main

import (
	"fmt"
)

func main() {
	var intSlice = []int{91, 42, 23, 14, 15, 76, 87, 28, 19, 95}
	chOdd := make(chan int)
	chEven := make(chan int)

	go odd(chOdd)
	go even(chEven)

	for _, i := range intSlice {
		if i%2 == 0 {
			chEven <- i
		} else {
			chOdd <- i
		}
	}

}

func odd(ch chan int) {
	for i := range ch {
		fmt.Println("Odd:", i)
	}
}

func even(ch chan int) {
	for i := range ch {
		fmt.Println("Even:", i)
	}
}
