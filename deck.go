package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Deck Of Cards
type deck []string

// Create a new deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// Iterate through the suits and values to create all permutations of cards
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" Of "+suit)
		}
	}

	// Return the new deck
	return cards
}

// Create a new deck from a loaded file
func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := string(byteSlice)
	c := strings.Split(s, ",")
	d := deck(c)
	return d

}

// Print all cards in the deck to the console
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i + 1, card)
	}
}

// Split the deck by the handSize and return two new decks
func (d deck) deal(handSize int) (deck, deck)  {
	return d[:handSize], d[handSize:]
}

// Shuffle all cards inside of the deck
func (d deck) shuffle() {

	// Create a truly random number generator by specifying a seed value to the rand package
	// To generate the randomized source int64, the UnixNano time function
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		// Get the new position
		newPosition := r.Intn(len(d) - 1)
		// Swap the elements
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// Save out a byte slice of the deck to the file system
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, d.toByteSlice(), 0666)
}

// Convert the deck slice to a single string
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// Convert the deck slice to a string, then to a byte slice
func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}


