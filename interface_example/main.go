package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct{}

// getGreeting implements Bot.
func (e EnglishBot) getGreeting() string {
	return "Hi..."
}

type SpanishBot struct{}

// getGreeting implements Bot.
func (s SpanishBot) getGreeting() string {
	return "Hola..."
}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}
	printMessage(eb)
	printMessage(sb)
}

func printMessage(b Bot) {
	fmt.Println(b.getGreeting())
}
