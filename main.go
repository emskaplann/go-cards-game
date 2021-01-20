package main

// func main() {
// 	// var card string = "Ace of Spades"
// 	// card := newCard()
// 	// card = "Five of Diamonds"
// 	cards := newDeck()
// 	hand, _ := deal(cards, 4)

// 	// hand.print()
// 	// remainingDeck.print()
// 	hand = hand.shuffle()
// 	hand.saveToFile("hand_save")

// 	handFromFile := newDeckFromFile("hand_save")
// 	handFromFile.print()
// }

func main() {
	newDeck := newDeck()
	newDeck = newDeck.shuffle()
	newDeck.print()
}
