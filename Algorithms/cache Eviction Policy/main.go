package main

import (
	"fmt"
)

// cachelist to store the result after every addition
type CacheList []map[int]string

type Stack struct {
	top      int
	mark     int
	capacity uint
	array    []interface{}
}

//Returns an initialized array
func (stack *Stack) Init(capacity uint, strategyID int) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.array = make([]interface{}, capacity)
	switch strategyID {
	case 1: // if LRU strategy, set mark to the starting
		stack.mark = 0
	case 2: //if MRU strategy, set mark to the end
		stack.mark = int(capacity) - 1
	}
	return stack
}

//Return a new stack
func NewStack(capacity uint, strategy int) *Stack {
	return new(Stack).Init(capacity, strategy)
}

//checks if the stack is full
func (stack *Stack) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

func (stack *Stack) Insert(strategyID int, data interface{}) string {
	var elem string
	switch strategyID {
	case 1:
		if stack.IsFull() {
			/* the "mark" is set at [0] for LRU by default,
			if the stack is full, then the "mark" will be
			shifted to the RHS after replacing the least recently used
			item
			*/
			stack.array[stack.mark] = data
			if stack.mark < int(stack.capacity) {
				stack.mark++
			} else {
				stack.mark = 0 // if mark reaches beyond end, set it to [0]
			}
		} else {
			stack.top++
			stack.array[stack.top] = data
		}
	case 2:
		if stack.IsFull() {

			/* the "mark" is set at [end] for MRU by default,
			if the stack is full, then the "mark" will be
			shifted to the LHS after replacing the most recently used
			item
			*/
			stack.array[stack.mark] = data
			if stack.mark == 0 {
				stack.mark = int(stack.capacity) - 1 // if mark reaches beyond [0], set it to [end]
			} else {
				stack.mark--
			}
		} else {
			stack.top++
			stack.array[stack.top] = data
		}
	default:
		fmt.Println("strategyID invalid ")
		return " "

	}
	for _, r := range stack.array {
		if r != nil {
			elem = elem + " " + fmt.Sprintf("%v", r)
		}
	}
	return elem

}

func (c CacheList) Get(key int) string {
	for _, r := range c {
		v, found := r[key] // if the key is found
		if found {
			return v
		}
	}

	return "Key Not Found"

}

// Returns the most recent sequence
func (c CacheList) StateOfCache() map[int]string {
	return c[len(c)-1]
}

// Helper method to take user input
func getInput(p string, t interface{}) interface{} {
	display(p)
	switch v := t.(type) {
	case int:
		var input int
		fmt.Scan(&input)
		return input
	case uint:
		var input uint
		fmt.Scan(&input)
		return input
	case string:
		var input string
		fmt.Scan(&input)
		return input
	default:
		fmt.Println("unknown type ", v)
	}
	return nil

}

//Helper method to print
func display(p ...interface{}) {
	fmt.Println()
	fmt.Println(p)
	fmt.Println()
}

func main() {
	var input int
	var cacheList CacheList
	cacheSize := getInput("Enter the cache size", uint(0)).(uint)
	strategyID := getInput("Select the strategy \n 1.LRU ( Least Recently Used) \n 2.MRU ( Most Recently Used) \n", 0).(int)
	// strategyID decides whether the CacheEviction is going to happen in LRU or MRU
	cache := NewStack(cacheSize, strategyID)

	for input != 9 { // 9 for program exit
		input = getInput("Select the operation to perform \n 1. Insert \n 2.Get \n 3.State of Cache \n(Press 9 to exit)", 0).(int)

		switch input {
		// user decides to insert
		case 1:
			key := getInput("please enter key and press enter", 0).(int) //converting the interface{} output of getInput method into int
			value := getInput("please enter value and press enter", "str").(string)
			sequence := cache.Insert(strategyID, value) //inserting according to the strategyID
			display("\nsequence is ", sequence)
			//storing the sequence after adding the value alongwith its key (provided by user)
			aMap := map[int]string{
				key: sequence,
			}
			//adding the resulting sequence to the list of results
			cacheList = append(cacheList, aMap)

			// user decides to use Get() method
		case 2:
			key := getInput("Enter the key", 0).(int)
			display(cacheList.Get(key))

			// user decided to use StateOfCache() method
		case 3:
			display(cacheList.StateOfCache())

			// user is fooling around
		default:
			display("enter a valid input \n")
		}

	}

}
