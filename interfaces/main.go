package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreating() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreating() string {
	return "Hello There!"
}
func (spanishBot) getGreating() string {
	return "Hola!"
}
func printGreeting(b bot) {
	fmt.Println(b.getGreating())
}
