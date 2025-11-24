package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

// Class = mehrere Schüler
type Class []Student

// Modules = modulnummer --> klassen (eine oder mehrere)
type Modules map[int][]Class

func main() {
	classA := Class{
		{FirstName: "Cornor", LastName: "McGregor"},
		{FirstName: "Jon", LastName: "Jones"},
		{FirstName: "Ilia", LastName: "Topuria"},
	}
	classB := Class{
		{FirstName: "Khabib", LastName: "Nurmagomedov"},
		{FirstName: "Khamzat", LastName: "Chimaev"},
		{FirstName: "Magomed ", LastName: "Ankalaev"},
	}

	modules := Modules{
		346: {classA, classB},
		320: {classA},
		114: {classB},
	}

	fmt.Println("Class A:", classA)
	fmt.Println("Class B:", classB)
	fmt.Println("Modules:", modules)
}
