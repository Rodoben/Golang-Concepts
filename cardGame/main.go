package main

func main() {

	// cards := newDeck()
	// cards.print()

	// // hand, remainingCards := deal(cards, 5)
	// // hand.print()
	// // remainingCards.print()
	// var n string
	// // cards := newDeck()
	// // fmt.Println(cards.toString())
	// fmt.Println("Enter the file name to be saved as")
	// fmt.Scanln(&n)
	// // cards.saveToFile(n)
	// cards := newDeckFromFile(n)
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
