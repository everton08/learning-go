package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("pepo_cards")

	// new deck from file
	// cards := newDeckFromFile("pepo_cards")
	// cards.print()
}
