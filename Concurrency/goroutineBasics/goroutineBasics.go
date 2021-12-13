package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string)
	fmt.Println("Hello World")
	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				ch <- "Goroutine: " + strconv.Itoa(i)
			}
		}(i)
	}

	for k := 1; k <= 100; k++ {
		fmt.Println(k, <-ch)
	}
}
