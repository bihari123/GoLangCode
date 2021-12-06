package isogram

import (
    "fmt"
    "strings"
    "unicode"
)



func IsIsogram(word string) bool {
    var sample =""

    for _,char:= range word{
       if string(char)==" " || string(char)=="-"{
           continue
       }
        if strings.Contains(sample,string(unicode.ToUpper(char))) || strings.Contains(sample,string(unicode.ToLower(char))){
            fmt.Printf("%c is already present",char)
            return false
        }
        sample+=string(char)
    }

    return true
	panic("Please implement the IsIsogram function")
}
