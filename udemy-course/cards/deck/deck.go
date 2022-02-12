package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func NewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) Print() {
	for i, v := range d {
		fmt.Printf("%d - %s \n", i+1, v)
	}
}

func Deal(d deck, handSize int8) (deck, deck) {
	return d[0:handSize], d[handSize:]
}

func (d deck) ToString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) SaveToFile(file string) {
	err := ioutil.WriteFile(file, []byte(d.ToString()), 0666)
	if err != nil {
		panic(err)
	}
	fmt.Printf("File saved to %s \n", file)
}

func NewDeckFromFile(file string) deck {
	bytesSlice, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("There was an error reading from file %s. Error: %+v \n", file, err)
		os.Exit(1)
	}

	stringSlice := strings.Split(string(bytesSlice), ",")

	return deck(stringSlice)
}

func (d deck) Shuffle() {
	seedSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seedSource)
	for i, _ := range d {
		newPos := r.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}
}
