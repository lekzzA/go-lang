package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create new type 'deck' of type slice of string.
type deck []string

// Error implements error.
func (d deck) Error() string {
	panic("unimplemented")
}

// Receiver function to print object of type deck.
func (d deck) print() {
	for i, val := range d {
		println("-", i, ": "+val)
	}
}

/**
* Return a new 'deck' of cards with 16 cards.
* With suites - Spades, Hearts, Diamonds and Clubs
* With numbers - Ace, Two, Three and Four
 */
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardNumbers := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardNumbers {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

/**
* Returns two decks. First one with number of cards of handSize, second remaining cards in deck.
 */
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/**
* Reciever method that saves the given deck into the provided filename.
 */
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

/**
* Receiver method to convert a slice of decks to a comma separated string.
 */
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

/**
* Returns a deck after loading it from the file filename.
 */
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		println("Error: ", err.Error())
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

/**
* Shuffles the given deck.
 */
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
