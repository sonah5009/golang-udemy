package main

import "fmt"

// deckSize := 20 // not allowed shortened delclare and assign :=
var deckSize int // requires `var`

// var deckSize = 20

func main() {
	deckSize = 20
	fmt.Println(deckSize)
}

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hi there!")
// }

// how to run the code
// In terminal
// go run main.go
