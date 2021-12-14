package purchase

import (
	"fmt"
	"sort"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind=="car"|| kind=="truck"{
		return true
	}
	return false
	panic("NeedsLicense not implemented")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	list:=[]string{}
	list=append(list,option1,option2)
	sort.Strings(list)
	return fmt.Sprintf("%s is clearly the better choice.",list[0])
	panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age <3{
		return 0.8*originalPrice
	} else if age<10 &&age>=3{
		return 0.7*originalPrice
	}else if age>=10{
	    return 0.5*originalPrice
	}
	panic("CalculateResellPrice not implemented")
}
