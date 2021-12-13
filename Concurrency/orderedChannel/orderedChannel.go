package main

import "fmt"

func sum(arr []int, ch chan int) {
	sum := 0
	for _, r := range arr {
		sum += r
	}
	ch <- sum
}

func main() {
	ch := make(chan int)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 12, 23, 34, 35, 56, 57}
	for i := 0; i < len(a); i += 4 {
		go sum(a[i:i+4], ch)
	}
	order := [4]int{}
	for i := 0; i < 4; i++ {
		order[i] = <-ch
	}
	close(ch)

	fmt.Println(order)

}
