package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("pepo_cards")

	// new deck from file
	// cards := newDeckFromFile("pepo_cards")
	// cards.print()
}
