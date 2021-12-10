package main

import "fmt"

// create a new type "deck"
// a slice of strings

type deck []string

func newDeck() deck {
	newDeck := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	jokers := []string{"Joker", "Joker", "Joker", "Joker"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newDeck = append(newDeck, value + " of "+ suit)
		}
	}

	for _, joker := range jokers {
		newDeck = append(newDeck, joker)
	}

	return newDeck
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

