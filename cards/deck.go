package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	d := deck{}
	cardSuits := []string{
		"Spades", "Diamonds", "Heart", "Clubs",
	}
	cardValues := []string{
		"Ace", "Two", "Three", "Four",
	}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, value+" of "+suit)
		}
	}
	return d
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(handsize int) (deck, deck) {
	return d[:handsize], d[:handsize]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveDeckToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[newPosition], d[i] = d[i], d[newPosition]
	}
}

func newDeckFromFile(filename string) (deck, error) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	str := string(bs)
	d := deck(strings.Split(str, ","))
	return d, nil
}
