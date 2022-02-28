package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newdeck() deck {
	cards := deck{}

	cardsuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardvalue := []string{"Ace", "one", "Two", "Three", "Four"}

	for _, suits := range cardsuits {
		for _, values := range cardvalue {
			cards = append(cards, values+" of "+suits)
		}
	}

	return cards

}
