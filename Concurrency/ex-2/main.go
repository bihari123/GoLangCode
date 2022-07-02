package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup){
	fmt.Println("started Goroutine, ",i) 

	time.Sleep(5*time.Second) 

	fmt.Println("Goroutine ",i," ended")
	wg.Done()
}

func main() {

	var wg  sync.WaitGroup 
 
 	for i:=0;i<4;i++{
 		wg.Add(1) 

 		go process(i,&wg)
 	}
wg.Wait()

fmt.Print("all Go routines fininshes")

}
