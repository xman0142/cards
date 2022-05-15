package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of 'deck'
//which is a slice of strings

type deck []string //Class Declaration

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)

		}
	}

	return cards

}

//d is a receiver value
//By creating a new type with a function that has a receiver, we are adding a 'method' to any value of that type
//d is cosidered self in Python and the deck class has a print method
//bunch of functions we can use with d
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

// d is receiver for the deck 'type', so we can call card.toString()
func (d deck) toString() string {
	//Type Conversion: Slice of type string deck
	//place the value that we want next to the value we do have
	return strings.Join(d, ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

//Return type of deck
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//Option #1 - log (print) the error and return a call to newDeck()
		//Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		//Exit(1) means to close down program
		os.Exit(1)

	}
	//String Conversion: Place the Value that we want to have i.e. string(), next to the value we do have i.e. bs (bytestring)
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// rand.Intn(len(d)-1) means generate a number between 0 and 1 less than the legnth of the slice (or d for deck)
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New((source))
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// Swap Element Logic: Take what ever is at new postion and assign it to d[i], then take whatever is at d[i] and assign to new postion
		d[i], d[newPosition] = d[newPosition], d[i]

	}

}
