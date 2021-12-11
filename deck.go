package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i, card := range d {
		ra := r.Intn(len(d) - 1)
		t := d[ra]
		d[ra] = card
		d[i] = t

		// or
		// d[i], d[r] = d[r], d[i]
	}
}