package main

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// 	hand, remainingcards := deal(cards, 5)

// 	hand.print()
// 	remainingcards.print()

// cards := newDeckFromFile("my_cards")
