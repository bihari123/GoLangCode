package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for i := range jobs {
		fmt.Println("worker ", id, " started the job ", i)
		time.Sleep(3 * time.Second)
		fmt.Println("worker ", id, " fimished the job ", i)
		result <- i * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for _ = range jobs {
		<-result
	}

}
