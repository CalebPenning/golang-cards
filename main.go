package main

import (
	// "container/list"
	"fmt"
)

func main() {
	cards := []string{"Ace of Diamonds", newCard()}

	for i := 0; i < len(cards); i++ {
		fmt.Println(cards[i])
	}
}

func newCard() string {
	return "Five of Diamonds"
}

// in go, arrays have a fixed size

// slices behave more like js arrays,
// can grow or shrink

