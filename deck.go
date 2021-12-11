package main

import (
	"fmt"
	"strings"
	"os"
	"math/rand"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	d := deck(s)
	return d
}

func (d deck) shuffle() {
	for i, card := range d {
		r := rand.Intn(len(d) - 1)
		t := d[r]
		d[r] = card
		d[i] = t
	}
}