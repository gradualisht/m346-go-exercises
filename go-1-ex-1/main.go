package main

import "fmt"

func main() {
    var firstName string = "Ylldrit"
    var lastName string = "Abazi"
    var dayOfBirth int = 17
    var monthOfBirth string = "August"
    var yearOfBirth int = 2008
    var numberOfSiblings int = 2
    var heightInMeters float64 = 1.80
    var zodiacSign string = "Löwe"

    fmt.Printf("Name: %s %s\n", firstName, lastName)
    fmt.Printf("Geburtsdatum: %d. %s %d\n", dayOfBirth, monthOfBirth, yearOfBirth)
    fmt.Printf("Geschwister: %d\n", numberOfSiblings)
    fmt.Printf("Grösse: %.2f m\n", heightInMeters)
    fmt.Printf("Sternzeichen: %s\n", zodiacSign)
}
