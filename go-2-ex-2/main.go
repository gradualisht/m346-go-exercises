package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	// direkt auffüllen bis index 4
	fibs[2] = fibs[0] + fibs[1] // 2
	fibs[3] = fibs[1] + fibs[2] // 3
	fibs[4] = fibs[2] + fibs[3] // 5

	// nächste zahl anhängen
	fibs = append(fibs, fibs[3]+fibs[4]) // 8

	// noch drei weitere anhängen
	fibs = append(fibs, fibs[len(fibs)-1]+fibs[len(fibs)-2]) // 13
	fibs = append(fibs, fibs[len(fibs)-1]+fibs[len(fibs)-2]) // 21
	fibs = append(fibs, fibs[len(fibs)-1]+fibs[len(fibs)-2]) // 34

	fmt.Println(fibs) // [1 1 2 3 5 8 13 21 34]
}
