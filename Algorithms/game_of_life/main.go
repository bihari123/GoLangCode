package main

import (
	"fmt"
	"game_of_life/game_board"
	"time"
)

func main(){
	var count int
	game:=game_board.New(10,10)
	fmt.Println("how many life cycles do you want to see?")
	fmt.Scanln(&count)
	for i:=0;i<count;i++{
		fmt.Printf("\n\ngenerating life cycle number: %d \n\n",i+1)
		time.Sleep(1000000000)
		game.PrintGrid()
		time.Sleep(100000)
		for j:=0;j<2;j++{
			fmt.Print("\n")
			time.Sleep(100000000)
		}

		for j:=0;j<100;j++{
			fmt.Print("-")
			time.Sleep(10000000)
		}

		for j:=0;j<2;j++{
			fmt.Print("\n")
			time.Sleep(100000000)
		}

		//Utility.ClearScreen()
		game.Update()

	}

	//game.TestPrintGrid()
}
