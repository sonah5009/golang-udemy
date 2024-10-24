package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// `:` for delcare (new variable), `=` redefine 
	card := "Ace of Spades" // compiler is gonna figure out the type
	card = "Five of Diamonds"
	fmt.Println(card)
}