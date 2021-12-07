package clock

import "fmt"

// Define the Clock type here.
type Clock struct{
    h,m int
}

func (c Clock)Reformat() (h,m int){
	for c.h<0 || c.h>=24 || c.m<0 || c.m>=60{

		if c.h<0 {
			c.h+=24
		}else if c.h>=24 {
			c.h-=24
		}
		if c.m<0 {
			c.h--
			c.m+=60
		}else if c.m>=60 {
			c.h++
			c.m-=60
		}
	}
	h, m = c.h, c.m
	return h,m
}

func New(h, m int) Clock {
	c:=Clock{h,m}
	c.h,c.m=c.Reformat()
	return c
	panic("Please implement the New function")
}

func (c Clock) Add(m int) Clock {
    // if the function returns pointers rather than values than the program wont work
    c.m+=m
	c.h,c.m=c.Reformat()
    return c
	panic("Please implement the Add function")
}

func (c Clock) Subtract(m int) Clock {
	c.m-=m
	c.h,c.m=c.Reformat()
	return c
	panic("Please implement the Subtract function")
}

func (c Clock) String() string {
	return  fmt.Sprintf("%.2d:%.2d",c.h,c.m)
	panic("Please implement the String function")
}
