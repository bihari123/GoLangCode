package gross

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    unitMap:=map[string]int{
        "quarter_of_a_dozen" : 3,
        "half_of_a_dozen"    : 6,
        "dozen"              : 12,
        "small_gross"        : 120,
        "gross"              : 144,
        "great_gross"        : 1728,
        }
    return unitMap    
	panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    newBill:=make(map[string]int)
    return newBill
	panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    
    for unitItemKey,_:=range units{
        if unit == unitItemKey{
            bill[item]+=units[unit]
            return true
        }
     }
     return false
       
	panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    flag1:=false
    flag2:=false
    fmt.Println("checking bill items")
    for billItemKey,_:=range bill{
        if item == billItemKey{
            flag1=true
            fmt.Println("bill item found")
        }
     }
     fmt.Println("checking unit items")
     for unitItemKey,_:=range units{
        if unit == unitItemKey{
            flag2=true
            fmt.Println("unit item found")
        }
     }
     if flag1 && flag2{
        if bill[item]-units[unit] < 0{
            return false
        }else if bill[item]-units[unit] == 0{
            delete(bill,item)
            return true
        }
         fmt.Printf("Before deduction %d \n",bill[item])
         bill[item]=bill[item]-units[unit]
         fmt.Printf("After deduction %d \n",bill[item])
     }
     return flag1 && flag2
       
   	panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    
    for billItemKey,_:=range bill{
        if item == billItemKey{
           return bill[billItemKey],true 
        }
     }
     
     return 0, false
       
	panic("Please implement the GetItem() function")
}
