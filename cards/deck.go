package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		// card := cardValues[i] + " of " + suit
		// cards = append(cards, card)
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
