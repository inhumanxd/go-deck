package main

import (
	"log"
	"os"
)

func main() {
	cards, err := newDeckFromFile("cards")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	println(len(cards))

	cardsInHand, _ := deal(cards, len(cards)/2)
	cardsInHand.shuffule()
	cardsInHand.print()

	// log.Println("cardsInHand")
	// cardsInHand.print()

	// log.Println("cardsInDeck")
	// cardsInDeck.print()

	cards.saveFile("cards")
}
