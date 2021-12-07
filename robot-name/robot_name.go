package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Define the Robot type here.
type Robot struct{
	name string
}
var nameList []string
var alphabet = make([]string, 0, 26)
var digit=make([]int,0,10)
// Loop over 'A' to 'Z' and keep on appending
// to alphabet slice
func init(){
	var ch byte
	n:="&5"
	nameList=append(nameList,n)

	for ch = 'A'; ch <= 'Z'; ch++ {
		alphabet = append(alphabet, string(ch))
	}
	for ch =0;ch<10;ch++{
		digit=append(digit,int(ch))
	}
}


//Generate count non-repeating random numbers ending with [start, end)
func generateRandomNumber (start int, end int, count int) [] int {
	//range check
	if end <start || (end-start) <count {
		return nil
	}

	//The result slice

	//Random number generator, add a timestamp to ensure that the random number generated each time is different
	r:= rand.New (rand.NewSource (time.Now (). UnixNano ()))

	nums :=  make ([] int, 0)
	for len (nums) <count {
		//Generate random numbers
		num:= r.Intn ((end-start)) + start

		//Check the weight
		exist:= false
		for _, v:= range nums {
			if v == num {
				exist = true
				break
			}
		}

		if! exist {
			nums = append (nums, num)
		}
	}

	return nums
}

func isContains(list []string,name string)bool{
	for _,elem:=range list{
		if strings.Compare(strings.Trim(elem," "),strings.Trim(name," "))==0{
			return true
		}
	}
	return false
}

func (r *Robot) Name() (string, error) {
    if len(r.name)>0{
		return r.name ,nil
	}
	r.name=""
	if len(nameList)>(26*26*10*10*10){ //checking if all possible name have been allocated using formula as studied in PnC for repeat cases
		return "",errors.New("all possible names already allocated")
	}
	for true{
		//fmt.Println("making name:   ",name)
		if len(r.name)<5{
			for i:=0;i<=1;i++{
				random:=generateRandomNumber(0,25,5)
				r.name=r.name[:len(r.name)]+string(alphabet[random[0]])
			}
			for i:=2;i<=4;i++{
				random:=generateRandomNumber(0,9,5)
				r.name=r.name[:len(r.name)]+strconv.Itoa(digit[random[0]])
			}
			if !isContains(nameList,r.name){
			   	break
			}
		}else{
			r.name=""
		}

	}
	//fmt.Println(name)

	nameList[len(nameList)-1]=r.name
	return r.name,nil
	panic("Please implement the Name function")
}

func (r *Robot) Reset() {
	if len(r.name)>0{
		for index,name:=range nameList{
			if name==r.name{
				fmt.Println("index: ",index)
				if index >0{
					nameList=append(nameList[0:index],nameList[index+1:]...)
				}
			}
		}
	}

	r.name=""

	//panic("Please implement the Reset function")
}
