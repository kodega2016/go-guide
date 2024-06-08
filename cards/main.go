package main

import (
	"fmt"
	"os"
)

func main() {
	d, err := newDeckFromFile("on_hands.txt")
	if err != nil {
		fmt.Println("files not found:", err)
		os.Exit(1)
	}
	d.print()
	// d := newDeck()
	// on_hands, _ := d.deal(4)
	// on_hands.shuffle()
	// err := on_hands.saveDeckToFile("on_hands.txt")
	// if err == nil {
	// 	fmt.Println("file saved successfully...")
	// }
}
