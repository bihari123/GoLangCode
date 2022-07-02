package main

import (
	"fmt"
	"time"
)

func process (ch chan string) {
  time.Sleep(11*time.Second)

  ch<-"process successful"
}

func main(){
  ch:=make (chan string)

  go process(ch)

  for {
    select{
    case v:=<-ch:
      fmt.Println("recieved the value ",v)
      return 
      default:
        fmt.Println("no value recieved")
    }
  }
}
