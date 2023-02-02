package main

import "fmt"

// Program Parameters
const players = 3    // Players per deck
const cardsDealt = 7 // Cards each deal

func main() {
	var player [players]deck
	cards := newDeck()
	cards.print()
	cards.shuffle()
	fmt.Println("\nCards are shuffled:")
	cards.print()

	for i := 0; i < players; i++ {
		player[i], cards = deal(cards, cardsDealt)
		fmt.Println("\nCards dealt for Player", (i + 1))
		player[i].print()
		fmt.Println("\nRemaining cards in deck:")
		cards.print()
	}

	// cards.print()

	//

	// fmt.Println("\nCards dealt Player 2")
	// hand.print()
	// cards.print()
	// cards.saveToFile("card_deck.txt")
}
