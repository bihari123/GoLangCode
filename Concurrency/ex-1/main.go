package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello goroutine")
	done <- true
}

func producer(chan1 chan int) {
	for i := 0; i < 10; i++ {
		chan1 <- i
	}
	close(chan1)
}

func write(ch chan int) {

	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Successfully wrote ", i, " to ch")
	}
	close(ch)
}

func main() {
	done := make(chan bool)
	go hello(done)

	<-done
	fmt.Println("main func")

	ch := make(chan int)

	go producer(ch)

	for {
		// if ok is false , if means that the channel has been closed and it will not send data anymore
		if v, ok := <-ch; ok {

			fmt.Println("Recieved : ", v)

		} else {
			break
		}

	}

	bufChan := make(chan int, 1)

	go write(bufChan)

	for v := range bufChan {
		fmt.Println("Read value :", v, " from channel")
		time.Sleep(5 * time.Second)
	}
}
