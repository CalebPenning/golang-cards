package main


func main() {
	cards := newDeckFromFile("my_cards")
	println(len(cards))
	println("HERE'S THE ORIGINAL DECK")
	cards.print()
	println("SHUFFLING")
	cards.shuffle()
	println("DONE SHUFFLING, HERE COMES RANDOM DECK")
	cards.print()

}