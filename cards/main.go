package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(card)
		// if we don't use 'i' inside => can see "declared and not used"
	}
}

func newCard() string {
	return "Five of Diamonds"
}
