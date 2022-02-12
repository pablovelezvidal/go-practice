package main

import (
	"cards/deck"
)

func main() {
	// cards := deck.NewDeck()

	// myHand, _ := deck.Deal(cards, 7)

	// myHand.SaveToFile("test.txt")

	myHand := deck.NewDeckFromFile("test.txt")

	// myHand.Print()

	myHand.Shuffle()

	myHand.Print()
}
