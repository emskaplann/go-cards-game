package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice(array) of strings
type deck []string

// (d deck): d means the copy of the variable we are using
// (d deck): deck means only variable type deck can access
// Go avoids any mention of 'this' and 'self'!
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) {
	deckInBytes := d.toBytes()
	err := ioutil.WriteFile(filename, deckInBytes, 0644)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func (d deck) toBytes() []byte {
	newString := []string(d)
	joinedString := strings.Join(newString, ",")
	return []byte(joinedString)
}

func newDeckFromFile(filename string) deck {
	deckInBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(deckInBytes), ","))
}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		randomIdx := r.Intn(len(d) - 1)
		// temp := d[randomIdx]
		// d[randomIdx] = card
		// d[i] = temp
		d[i], d[randomIdx] = d[randomIdx], d[i] // cool one-line swap
	}
	return d
}
