package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting new card program...")
	cards := newDeck()
	cards.print()
	fmt.Println("Done!")
}

// in go, arrays have a fixed size

// slices behave more like js arrays,
// can grow or shrink
