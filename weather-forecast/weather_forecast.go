// Package weather: Package containing implementation of the weather forecast logic.
package weather

// CurrentCondition: Variable storing the current weather condition.
var CurrentCondition string

// CurrentLocation: Variable storing the location for which the weather is being forcasted.
var CurrentLocation string

// Forecast: Function that outputs the weather forecast for a given input city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
