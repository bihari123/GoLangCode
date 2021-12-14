package scrabble

import "unicode"

type Group []rune
var LetterValue=[] struct{
    Group
    value int
} {
    {Group{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T' },1},
    {Group{'D', 'G'},2},
    {Group{'B', 'C', 'M', 'P'},3},
    {Group{'F', 'H', 'V', 'W', 'Y'},4},
    {Group{'K'},5},
    {Group{'J', 'X'},8},
    {Group{'Q', 'Z' },10},
}

func Score(word string) int {
    score:=0
    for _,charWord:= range word{
        for index,group:= range LetterValue{
            for _,charGroup:= range group.Group{
                if charGroup ==unicode.ToUpper(charWord){
                    score+=LetterValue[index].value
                }
            }
        }
    }
    return score
    // panic("Please implement the Score function")
}
