package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.suffle()
	cards.print()
}