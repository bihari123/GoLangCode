package secret

import (
	"fmt"
	"strconv"
)

var instruction = map[int]string{
	1:    "wink",
	10:   "double blink",
	100:  "close your eyes",
	1000: "jump",
}

func revStr(str string) (result string) {
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return
}

func revArr(arr []int) []int {
	var newArr []int
	for i := len(arr) - 1; i >= 0; i-- {
		newArr = append(newArr, arr[i])
	}
	return newArr
}

func binToComm(num int) []int {
	var arr []int
	divisor := 10

	for divisor <= (num*10) || divisor == 10 {
		elem := num % divisor
		if elem > 0 {
			arr = append(arr, elem)
			num -= elem
		}
		divisor *= 10
	}
	return arr
}

func decToBin(d int) int {
	bStr := ""
	for d > 0 {
		bStr += fmt.Sprintf("%d", (d % 2))
		d /= 2
	}
	bStr = revStr(bStr)
	b, _ := strconv.Atoi(bStr)
	return b

}

func Handshake(code uint) []string {
	var inst []string
	if code > 0 {
		arr := binToComm(decToBin(int(code)))

		if arr[len(arr)-1] == 10000 {
			arr = arr[0:(len(arr) - 1)]
			arr = revArr(arr)
		}

		for _, elem := range arr {
			inst = append(inst, instruction[elem])
		}

	}

	return inst
	panic("Please implement the Handshake function")
}
