package main

import "fmt"

func main() {
    miles := 10.0
    fahrenheit := 68.0

    kilometers := miles * 1.60934
    celsius := (fahrenheit - 32) * 5 / 9

    fmt.Printf("%.2f Meilen = %.2f km\n", miles, kilometers)
    fmt.Printf("%.2f °F = %.2f °C\n", fahrenheit, celsius)


    kilometers2 := 16.0934
    celsius2 := 20.0

    miles2 := kilometers2 / 1.60934
    fahrenheit2 := (celsius2 * 9 / 5) + 32

    fmt.Printf("%.2f km = %.2f Meilen\n", kilometers2, miles2)
    fmt.Printf("%.2f °C = %.2f °F\n", celsius2, fahrenheit2)
}
