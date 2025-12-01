package main

import "fmt"

// typdefinitionen
type Celsius float64
type Fahrenheit float64

// funktionen für umrechnung
func convertCelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}
func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// methoden auf typen
func (c Celsius) ConvertToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func (f Fahrenheit) ConvertToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	// funktionsaufrufe
	fmt.Println(convertCelsiusToFahrenheit(0))   
	fmt.Println(convertCelsiusToFahrenheit(23))  
	fmt.Println(convertCelsiusToFahrenheit(100)) 

	// rückrechnung
	fmt.Println(convertFahrenheitToCelsius(32))  
	fmt.Println(convertFahrenheitToCelsius(73.4))
	fmt.Println(convertFahrenheitToCelsius(212)) 

	// methodenaufrufe
	var cozy Celsius = 23
	fmt.Println(cozy.ConvertToFahrenheit())

	var cold Fahrenheit = 15.3
	fmt.Println(cold.ConvertToCelsius()) 
}
