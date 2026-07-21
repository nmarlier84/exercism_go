// Package weather provides tools to define weather forecast.
package weather

var (
	//CurrentCondition represents the current weather.
	CurrentCondition string
	//CurrentLocation represents the current location.
	CurrentLocation string
)

// Forecast returns the current condition in a specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
