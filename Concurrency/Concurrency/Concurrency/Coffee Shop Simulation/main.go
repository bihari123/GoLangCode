package main

import (
	"fmt"
	"sync"
	"time"
)

type Ingredients map[string]int

type CoffeeMachine struct {
	outlets int
	Ingredients
	sync.Mutex
}

func (c *CoffeeMachine) New(outlets int) *CoffeeMachine {
	c.outlets = outlets
	c.Ingredients = make(Ingredients)
	c.Ingredients["water"] = 1000
	c.Ingredients["milk"] = 1000
	c.Ingredients["tea_leaves"] = 100
	c.Ingredients["coffee_syrup"] = 100
	c.Ingredients["ginger_syrup"] = 100
	c.Ingredients["sugar_syrup"] = 100
	c.Ingredients["elaichi_syrup"] = 100
	return c
}

/*func showAvailable(options []map[string]int, i Ingredients) []map[string]int {
	var result []map[string]int
	for j, _ := range options {
		for k, v := range options[j] {
			if i[k] > v {
				result = append(result, options[j])
			}
		}
	}

	return result
}
*/

func (m *CoffeeMachine) IsStock(i Ingredients) bool {
	flag := true
	for k, v := range i {
		if v > m.Ingredients[k] {
			flag = false
		}
	}
	return flag
}
func (m *CoffeeMachine) Refill(lowstock []string) {
	for _, v := range lowstock {
		if v == "water" || v == "milk" {
			m.Ingredients[v] = 1000
		} else {
			m.Ingredients[v] = 100
		}
	}
}
func (m Ingredients) LowOnStock() []string {
	var result []string
	for k, v := range m {
		if v < 50 {
			result = append(result, k)
		}
	}
	return result
}

func (m *CoffeeMachine) MakeDrink(i Ingredients) {
	m.Lock()
	defer m.Unlock()
	for k, v := range i {
		m.Ingredients[k] -= v
	}
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
	fmt.Println(p...)
	fmt.Println()
}

func main() {
	ginger := Ingredients{
		"water":        50,
		"milk":         10,
		"tea_leaves":   10,
		"ginger_syrup": 5,
		"sugar_syrup":  10,
	}

	elaichi := Ingredients{
		"water":         50,
		"milk":          10,
		"tea_leaves":    10,
		"elaichi_syrup": 5,
		"sugar_syrup":   10,
	}

	coffee := Ingredients{
		"water":        50,
		"milk":         10,
		"coffee_syrup": 10,
		"sugar_syrup":  10,
	}

	hot_milk := Ingredients{
		"milk": 10,
	}

	hot_water := Ingredients{
		"water": 50,
	}

	outlets := getInput("Enter the number of outlets", 0).(int)

	machine := new(CoffeeMachine).New(outlets)
	machine_Ingredients := machine.Ingredients
	options := []map[string]int{
		ginger,
		elaichi,
		coffee,
		hot_milk,
		hot_water,
	}
	var choice int
	for choice != 9 {
		display("Place your order.(Press 9 to exit)\nthe following drinks are available \n 1. ginger \n 2. elaichi \n 3. coffee \n 4. hot milk\n 5. Hot Water")

		choice = getInput("which one do you like?", 0).(int)

		if choice == 9 {
			break
		}

		display(choice)
		time.Sleep(2000)
		drink := options[choice-1]
		if machine.IsStock(drink) {

			machine.MakeDrink(drink)

		} else {
			display("not enough ingredinets,refill required")
			time.Sleep(2000)

		}

		display(machine_Ingredients)
		time.Sleep(2000)
		lowStock := machine_Ingredients.LowOnStock()
		if len(lowStock) > 0 {
			display(" the following items are low on stock \n", machine_Ingredients.LowOnStock())
			input := getInput("would you like to refill? (Press 1 for yes and 0 for no", 0).(int)
			switch input {
			case 1:
				machine.Refill(lowStock)
				display(machine_Ingredients)
			case 0:

				display("Make your next order")
			}
			time.Sleep(1000)
		}

	}

}
