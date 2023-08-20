package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Joker", "Queen", "King"}

	for suitsIndex := 0; suitsIndex < len(cardSuits); suitsIndex++ {
		for valIndex := 0; valIndex < len(cardValues); valIndex++ {
			cards = append(cards, cardValues[valIndex]+" Of "+cardSuits[suitsIndex])
		}
	}

	return cards
}

func (d deck) print() {
	for _, v := range d {
		log.Println(v)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	bs, err := os.ReadFile(filename)

	if err != nil {
		log.Println("Error: ", err)
		return newDeck(), nil
	}

	cardsDeck := deck(strings.Split(string(bs), ","))

	return cardsDeck, err
}

func (d deck) shuffule() deck {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	for i := range d {
		swapPosition := r.Intn(len(d) - 1)

		// temp := v

		// d[i] = d[swapPosition]
		// d[swapPosition] = temp

		d[i], d[swapPosition] = d[swapPosition], d[i]
	}

	return d
}
