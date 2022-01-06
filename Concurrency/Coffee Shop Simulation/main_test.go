package main

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

var ginger = Ingredients{
	"water":        50,
	"milk":         10,
	"tea_leaves":   10,
	"ginger_syrup": 5,
	"sugar_syrup":  10,
}

var elaichi = Ingredients{
	"water":         50,
	"milk":          10,
	"tea_leaves":    10,
	"elaichi_syrup": 5,
	"sugar_syrup":   10,
}

var coffee = Ingredients{
	"water":        50,
	"milk":         10,
	"coffee_syrup": 10,
	"sugar_syrup":  10,
}

var hot_milk = Ingredients{
	"milk": 10,
}

var hot_water = Ingredients{
	"water": 50,
}

var outlets = 5

var machine = new(CoffeeMachine).New(outlets)
var machine_ingredients = machine.Ingredients
var options = []map[string]int{
	ginger,
	elaichi,
	coffee,
	hot_milk,
	hot_water,
}

func TestRefill(t *testing.T) {
	machine_ingredients["water"] = 40
	expectedOutput := 1000
	machine.Refill([]string{"water"})
	output := machine_ingredients["water"]
	if output != expectedOutput {
		t.Errorf("Failed ! got %v want %c", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}

}
func TestLowOnStock(t *testing.T) {
	machine_ingredients["water"] = 40
	if len(machine.LowOnStock()) > 0 {
		t.Logf("Success !")
	} else {
		t.Errorf("Failed ! got %v want %c", len(machine.LowOnStock()), 1)
	}
}
func TestIsStock(t *testing.T) {
	machine_ingredients["milk"] = 40
	if machine.IsStock(ginger) {
		t.Errorf("got true, expected false")
	} else {
		t.Logf("Success !")
	}
}

func TestConcDrinks(t *testing.T) {
	if runtime.NumCPU() < 2 {
		//t.Skip("Multiple CPU cores required for concurrency tests.")
	}
	if runtime.GOMAXPROCS(0) < 2 {
		runtime.GOMAXPROCS(2)
	}
	newOutlets := 5
	var m = new(CoffeeMachine).New(newOutlets)
	if m == nil {
		t.Fatal("new(CoffeeMachine).New(outlets) = nil, want non-nil *CoffeMachine.")
	}
	m_ingredients := m.Ingredients

	var start, g sync.WaitGroup
	start.Add(1)
	g.Add(newOutlets)

	for i := 0; i < newOutlets; i++ {
		go func() { // deposit
			start.Wait()
			//a.Deposit(amt) // ignore return values

			choice := rand.Intn(4 - 0)

			drink := options[choice-1]
			if machine.IsStock(drink) {

				machine.MakeDrink(drink)

			} else {
				//	display("not enough ingredinets,refill required")

			}

			//display(machine_ingredients)
			lowStock := m_ingredients.LowOnStock()
			if len(lowStock) > 0 {
				//	display(" the following items are low on stock \n", machine_ingredients.lowOnStock())
				input := 1
				switch input {
				case 1:
					machine.Refill(lowStock)
					//display(machine_ingredients)
				case 0:

					//display("Make your next order")
				}
			}

			g.Done()
		}()

	}
	start.Done()
	g.Wait()

	if len(m_ingredients.LowOnStock()) > 0 {
		t.Errorf("item on low stock, expected none")
	}
}
