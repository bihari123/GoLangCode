package main

import (
"fmt"
"math/rand"
"time"
)


//Generate count non-repeating random numbers ending with [start, end)
func generateRandomNumber (start int, end int, count int) [] int {
	//range check
	if end <start || (end-start) <count {
		return nil
	}

	//The result slice

	//Random number generator, add a timestamp to ensure that the random number generated each time is different
	r:= rand.New (rand.NewSource (time.Now (). UnixNano ()))

	nums :=  make ([] int, 0)
	for len (nums) <count {
		//Generate random numbers
		num:= r.Intn ((end-start)) + start

		//Check the weight
		exist:= false
		for _, v:= range nums {
			if v == num {
				exist = true
				break
			}
		}

		if! exist {
			nums = append (nums, num)
		}
	}

	return nums
}

func main(){
	fmt.Println(generateRandomNumber (0, 280, 17))
}
