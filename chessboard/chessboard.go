package chessboard

import "fmt"

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool
// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank 

func boolToStr(i bool) string{
    str:=""
	if i{
		str="true"
	}else{
		str="false"
	}
	return str
}


// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
    count:=0
    for _,occupy:= range cb[rank]{
        if occupy{
            count++
        }
    }
    return count
	panic("Please implement CountInRank()")
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
    count:=0
	if file >=1 && file<=8{
		for index,_:= range cb{
			fmt.Printf("checking.... rank: %s \n",index)
			if cb[index][file - 1]{
				fmt.Printf("rank: %s, file: %d, occupied: %s \n",index,file,boolToStr(cb[index][file-1]))
				count++
			}else{
				fmt.Printf("nothing found at rank : %s, file: %d \n",index,file)
			}
		}
	}

    return count
	panic("Please implement CountInFile()")
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count:=0
	for rank,_:= range cb{
	    for  range cb[rank]{
			count++
     	}


	}
	return count
	panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
    count:=0
	for rank,_:= range cb{
	    for _,file:= range cb[rank]{
	     	if file{
		        count++
	        }
	   }
	}
	return count
	panic("Please implement CountOccupied()")
}
