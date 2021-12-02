package lasagna

import "fmt"

// TODO: define the 'OvenTime' constant
 var OvenTime int=40
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	fmt.Printf("remaining time: %d \n", OvenTime - actualMinutesInOven)
	//panic("RemainingOvenTime not implemented")

	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	//panic("PreparationTime not implemented")
	return numberOfLayers * 2
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	//panic("ElapsedTime not implemented")
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}


func main() {
	fmt.Printf(" the preparation time is: %d", PreparationTime(2))
	//fmt.Printf(" the remaining time is: %d \n", RemainingOvenTime(25))
}

