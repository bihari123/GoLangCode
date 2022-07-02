package cars

const CarPerHour=221
// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed>=0 &&speed <1{
		return 0
	}
	if speed<=4 &&speed >=1{
		return 1
	}else if speed<=8{
		return 0.9
	} else if speed <=10{
		return 0.77
	}
	panic("SuccessRate not implemented - enter a valid value")
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return float64(CarPerHour * speed) * SuccessRate(speed)
	panic("CalculateProductionRatePerHour not implemented")
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed)/60)
	panic("CalculateProductionRatePerMinute not implemented")
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	production:=CalculateProductionRatePerHour(speed)
	if production>limit{
		return limit
	}else{
		return production
	}
	panic("CalculateLimitedProductionRatePerHour not implemented")
}
