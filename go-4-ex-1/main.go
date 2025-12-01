package main

import (
	"errors"
	"fmt"
)

// computeGrade berechnet eine Note von 1–6 anhand der erreichten Punkte
// formel: 1 + 5 * (gotPoints / maxPoints)
// gibt zusätzlich einen Fehler zurück, falls Eingaben unlogisch sind
func computeGrade(gotPoints, maxPoints float64) (float64, error) {
	if maxPoints <= 0 {
		return 0, errors.New("maxPoints muss > 0 sein")
	}
	if gotPoints < 0 || gotPoints > maxPoints {
		return 0, errors.New("gotPoints ist ausserhalb des gültigen Bereichs")
	}

	grade := 1 + 5*(gotPoints/maxPoints)
	return grade, nil
}

func main() {
	// 3 Tests mit unterschiedlichen Punktzahlen
	results := [][2]float64{
		{17.5, 28.0},
		{24.0, 30.0},
		{30.0, 30.0}, 
	}

	for _, r := range results {
		if grade, err := computeGrade(r[0], r[1]); err == nil {
			fmt.Printf("%.1f von %.1f --> Note: %.3f\n", r[0], r[1], grade)
		} else {
			fmt.Println("Fehler:", err)
		}
	}
}
