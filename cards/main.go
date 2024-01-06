package main

func main() {
	cards := createDeckFromFile("hello")
	cards.shuffle()
	cards.print()
}
