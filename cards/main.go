package main

import "fmt"

func main() {

	cards := []string{newcards(), "spades of five"}

	// to append value to a slice
	cards = append(cards, "three of hearts")
	for i, card := range cards {
		fmt.Println(i, card)
	}
	// fmt.Println(cards)

}

func newcards() string {
	return "ace of diamonds"
}
