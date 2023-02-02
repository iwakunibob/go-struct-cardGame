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
	cards := deck{}
	cardRanks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	cardSuits := []rune{'♦', '♥', '♠', '♣'}
	for _, cr := range cardRanks {
		for _, cs := range cardSuits {
			card := cr + string(cs)
			cards = append(cards, card)
		}
	}
	return cards
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Printf("%s ", card)
		if ((i % 12) == 11) || (i >= (len(d) - 1)) {
			fmt.Println()
		} else if (i % 4) == 3 {
			fmt.Print(" ")
		}
	}
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(filename string) error {
	outstring := d.toString()
	return os.WriteFile(filename, []byte(outstring), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 log the error and return a call to newDeck()
		// Option #2 log the error and quit program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	arystr := strings.Split(string(byteSlice), ", ")
	return deck(arystr)
}
