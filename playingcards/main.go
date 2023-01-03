package main

// import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	// cards := newDeckFromFile("mycards")
	// cards.saveToFile("mycards")

	// fmt.Println(cards.toString())
	cards.print()

}
