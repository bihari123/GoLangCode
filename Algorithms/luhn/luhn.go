package luhn

import (
    "fmt"
    "strings"
)

func modify(num uint8) int{

    if ((int(num)-int('0'))*2) >9{
        return ((int(num)-int('0'))*2) -9

    }else{
       return (int(num)-int('0'))*2
    }
}

func Valid(id string) bool {
    sum:=0
    id=strings.ReplaceAll(id," ","")
    fmt.Printf("\n\n")
    if len(id)<=1{
        return false
    }
    count:=1
    for i:=len(id)-1;i>=0;i--{
        num:=id[i]

        // (int(num)-int('0'))<0 ||(int(num)-int('0'))>9 for eliminating speacial characters
        if  (int(num)-int('0'))<0 ||(int(num)-int('0'))>9{
            return false
        }
        if count%2 == 0{
            fmt.Printf("Phase %d....Adding %c to sum %d \n",i,num,sum)
            sum+=modify(num)
            fmt.Printf("Interval Test %d.....the sum is %d \n",i,sum)
        }else{
            fmt.Printf("Phase %d....Adding %c to sum %d \n",i,num,sum)
            sum+=(int(num)-int('0'))
            fmt.Printf("Interval Test %d.....the sum is %d \n",i,sum)

        }
        count++
    }
    fmt.Printf("the sum is %d \n",sum)
  
    return sum %10 ==0
	panic("Please implement the Valid function")
}
