package say

import (
	"fmt"
	"strings"
)

var digit = map[int]string{
	0: "",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var digitPlace = map[int]string{
	1: "",
	2: "ty",
	3: "hundred",
	4: "thousand",
	5: "million",
	6: "billion",
	7: "trillion",
}

var digitAtTenth = map[int]string{
	2: "twen",
	3: "thir",
	4: "for",
	5: "fif",
	6: "six",
	8: "eigh",
	9: "nine",
}

var misc = map[int]string{
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}


var mark []int64
var list []int64

func Say(n int64) (string, bool) {
	var length = 0
	if n > 999999999999 || n < ((-1)*999999999999) {
		return "Number out of range", false
	}

	flag := true
	result := ""

	if n < 0 {
		n = n * (-1)
		return "",false
		result += "minus "
	}



	for n > 0 {


		o := n % 10
		t := ((n % 100) - o) / 10
		h := ((n % 1000) - (t * 10) - o) / 100
		var ones, tenth, hundredth string
		ones = digit[int(o)]


		if t > 0 {

			if t != 1 {
				tenth = digitAtTenth[int(t)] + digitPlace[2]

			} else {
				index := (t * 10) + o
				ones = ""
				tenth = misc[int(index)]
			}
		} else {
			tenth = ""
		}

		if h > 0 {

			hundredth = digit[int(h)] +" "+ digitPlace[3]
		} else {
			hundredth = ""
		}

		if length < 3 {
			if tenth!= "" && t != 1 && ones!="" {
				result = fmt.Sprintf("%s %s-%s ", hundredth, tenth, ones)
			}else {
				result = fmt.Sprintf("%s %s %s ", hundredth, tenth, ones)
			}

		} else {
			place := (length / 3) + 3
			if o>0 || t>0 || h>0{
				if tenth!= "" && t != 1 && ones!="" {
					result = strings.TrimSpace(fmt.Sprintf("%s %s-%s ", hundredth, tenth, ones) )+" "+ digitPlace[place] +" "+ result
				}else {
					result = strings.TrimSpace(fmt.Sprintf("%s %s %s ", hundredth, tenth, ones) )+" "+ digitPlace[place] +" "+ result
				}

			}

		}

		n = n / 1000
		length += 3

	}


    if result == ""{
		result="zero"
	}
	return strings.TrimSpace(result), flag
	panic("Please implement the Say function")
}
