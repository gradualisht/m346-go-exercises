package main

import (
	"fmt"
	"math"
)

// computeDiscriminant berechnet die D
func computeDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

// computeQuadraticFormula liefert Lösungen mit gleichgung
// anhand der Mitternachtsformel.
func computeQuadraticFormula(a, b, c float64) []float64 {
	D := computeDiscriminant(a, b, c)
	if D > 0 {
		return []float64{
			(-b + math.Sqrt(D)) / (2 * a),
			(-b - math.Sqrt(D)) / (2 * a),
		}
	} else if D == 0 {
		return []float64{(-b) / (2 * a)}
	}
	return []float64{} // keine Lösung (komplex)
}

func main() {
	fmt.Println(computeQuadraticFormula(3, 4, 1)) 
	fmt.Println(computeQuadraticFormula(2, 4, 2)) 
	fmt.Println(computeQuadraticFormula(3, 4, 2))

	// diskriminanten-tests
	fmt.Println(computeDiscriminant(3, 4, 1))
	fmt.Println(computeDiscriminant(2, 4, 2)) 
	fmt.Println(computeDiscriminant(3, 4, 2)) 
}
