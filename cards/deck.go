package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	symbols := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardsValue := []string{"Ace", "Two", "Three", "Four"}

	for _, sym := range symbols {
		for _, val := range cardsValue {
			card := sym + " of " + val
			cards = append(cards, card)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(fileName string) error {
	stringDeck := strings.Join(d, ",")
	byteDeck := []byte(stringDeck)

	err := ioutil.WriteFile(fileName, byteDeck, 0644)
	return err
}

func newDeckFromFile(fileName string) deck {
	byteDeck, error := ioutil.ReadFile(fileName)

	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	s := strings.Split(string(byteDeck), ",")

	return deck(s)

}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
