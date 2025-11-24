package main

import "fmt"

const (
	Diamonds = '\u25c6' // ◆
	Spades   = '\u2660' // ♠
	Clubs    = '\u2663' // ♣
	Hearts   = '\u2665' // ♥
	Six      = '\u2465'
	Seven    = '\u2466'
	Eight    = '\u2467'
	Nine     = '\u2468'
	Ten      = '\u2469'
	Jack     = 'J'
	Queen    = 'Q'
	King     = 'K'
	Ace      = 'A'
)

func main() {
	suits := []rune{Diamonds, Spades, Clubs, Hearts}
	ranks := []rune{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	for _, rank := range ranks {
		for _, suit := range suits {
			fmt.Printf("%c%c\t", suit, rank)
		}
		fmt.Println()
	}
}