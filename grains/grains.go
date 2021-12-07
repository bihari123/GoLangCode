package grains

import (
	"errors"
	"math"
)
var total uint64=0  // declared total as uint64 to avoid oerflow
var squares =make(map[int]uint64)
func init(){
    for i:=1;i<=64;i++{
	    	squares[i]=uint64(math.Pow(2, float64(i-1)))
			total+=squares[i]
    }
}
func Square(number int) (uint64, error) {
    if number>0 && number <=64{
        return squares[number], nil
    }
    return 0,errors.New("enter a valid number")
	panic("Please implement the Square function")
}

func Total() uint64 {
	return total
	//return result
   panic("Please implement the Total function")
}
