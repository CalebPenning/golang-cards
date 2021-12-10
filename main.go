package main

import (
	// "container/list"
	"fmt"
)

func main() {
	cards := deck{"Ace of Diamonds", newCard()}

	// classic explicitly declared loop
	for i := 0; i < len(cards); i++ {
		fmt.Println(cards[i])
	}

	// // alternate syntax
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
}

func newCard() string {
	return "Five of Diamonds"
}

// in go, arrays have a fixed size

// slices behave more like js arrays,
// can grow or shrink
