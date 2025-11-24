package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

// BirthDate hält das Geburtsdatum
type BirthDate struct {
	DayOfBirth   byte
	MonthOfBirth byte
	YearOfBirth  int16
}

type Profile struct {
	// Einbetten (promoted fields), wie im Theorie-Teil gezeigt
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName: FullName{
			FirstName: "Ylldrit",
			LastName:  "Abazi",
		},
		BirthDate: BirthDate{
			DayOfBirth:   17,
			MonthOfBirth: 8,
			YearOfBirth:  2008,
		},
		NumberOfSiblings: 2,
		ZodiacSign:       '\u264C', // ♌ Löwe
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// du bekommst ein weiteres Geschwister --> +1
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
