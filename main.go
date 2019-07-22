package main

func main() {

	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// type conversion
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// // fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my")
	cards.print()

}
