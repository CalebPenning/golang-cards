package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting...")
	cards := newDeck()
	cards.saveToFile("my_cards")
	fmt.Println("Finished saving file")
}