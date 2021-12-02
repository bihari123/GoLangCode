// Package weather contain tolls for weather forecast.
package weather

// CurrentCondition describes current condition.
var CurrentCondition string

// CurrentLocation describes the current location.
var CurrentLocation string

// Forecast returns current location with current conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
