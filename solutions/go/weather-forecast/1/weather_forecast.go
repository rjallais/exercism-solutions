// Package weather provides a tool to show a location's forecast.
package weather

// CurrentCondition represents a weather condition in the current time.
var CurrentCondition string

// CurrentLocation represents a city in which certain weather condition is occurring.
var CurrentLocation string

// Forecast returns a string containing a city and describing its weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
