package main

import "fmt"

func main() {
	c := generate()
	recieve(c)
}

func generate() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func recieve(c <-chan int) {
	for i := range c {
		fmt.Println(i, <-c)
	}
}
