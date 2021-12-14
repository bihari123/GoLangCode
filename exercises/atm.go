package main
import (
       "fmt"
       
       
        )



func main(){
	// your code goes here
 	var amount int
	var balance float64
	fmt.Scanln(&amount,&balance)
	//fmt.Scanln(&balance)
    if float64(amount) + 0.5 <= balance  && amount % 5 == 0{
	    
	    balance=balance - float64(amount) - 0.5 
	}
	    
	fmt.Printf("%0.2f",balance)
	
	
	
}



