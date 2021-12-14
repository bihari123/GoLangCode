package encode

import (
	"fmt"
	"strconv"
	"unicode"
)

func RunLengthEncode(input string) string {
	rep := 1
	var result string
	for i, _ := range input {
		if i+1 < len(input) {
			if input[i] == input[i+1] {
				rep++
			} else {
				repStr:=""
				if rep>1{
					repStr = strconv.Itoa(rep)
				}
				result=result+repStr+string(input[i])
				rep=1
			}

		}else{
			if rep==1{
				result+=string(input[i])
			}else{
				repStr := strconv.Itoa(rep)
				result=result+repStr+string(input[i])
			}
		}
	}
	return result
	panic("Please implement the RunLengthEncode function")
}
func toInt(str string) int{
	n:=0
	for i:=0;i<len(str);i++{
		n=n*10 + (int(str[i]) -48)
	}
	return n
}

func RunLengthDecode(input string) string {
	var result string
	for i:=0;i<len(input);i++{
		if unicode.IsDigit(rune(input[i])){
			var repStr string
			for unicode.IsDigit(rune(input[i])){
				 repStr+=string(input[i])
				 i=i+1
			}
			rep:=toInt(repStr)
            fmt.Println("rep: ",rep,"i: ",i)
			for j:=0;j<rep;j++{
				result+=string(input[i])
			}

		}else{
			result+=string(input[i])
		}
	}
	return result
	panic("Please implement the RunLengthDecode function")
}
