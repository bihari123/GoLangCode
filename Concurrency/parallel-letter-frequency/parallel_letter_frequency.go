package letter

import (
	"fmt"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	fmt.Printf("running goroutine for %s \n",s)
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {

	//messages := make(chan string)
	m := FreqMap{}
	ch:=make(chan FreqMap,len(l))
	for _, s := range l {
    	go func(s string) {

			ch <- Frequency(s) // first we have to write to the channel before we read from it
			                   // if we read from it before writing, then it will give DeadLock and
							   //program will panic

		}(s)

	}
	for  range l{

		if len(m)>0{
			newFreq:= <-ch
			for char,count:= range newFreq{
				m[char]+=count
			}

		}else{
			m=<-ch

		}
	}
return m

}
