// Package weather provides weather forecasting for Goblinocus.
package weather

// CurrentCondition represents the current conditions.
var CurrentCondition string

// CurrentLocation represents the current location.
var CurrentLocation string

// Forecast provices a message for the current conditions at provided city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
