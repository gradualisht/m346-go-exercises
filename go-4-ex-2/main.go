package main

import (
	"fmt"
	"math"
)

// computeHypotenuse berechnet Hypotenuse c via formel
func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

// truct-Typ 
type ShortSides struct {
	A float64
	B float64
}

// hypotenuse ist methode 
func (s ShortSides) Hypotenuse() float64 {
	return math.Sqrt(math.Pow(s.A, 2) + math.Pow(s.B, 2))
}

func main() {
	// funktionstests
	fmt.Println(computeHypotenuse(3, 4))
	fmt.Println(computeHypotenuse(5, 12)) 
	fmt.Println(computeHypotenuse(8, 15)) 

	// methoden-tests
	pairs := []ShortSides{{3, 4}, {5, 12}, {8, 15}}
	for _, p := range pairs {
		fmt.Printf("(%.0f, %.0f) --> %.1f\n", p.A, p.B, p.Hypotenuse())
	}
}
