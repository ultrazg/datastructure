package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Value string
}

var suits = []string{"Clubs", "Diamonds", "Hearts", "Spades"}
var values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

func main() {
	rand.Seed(time.Now().UnixNano())

	deck := generateDeck()
	shuffledDeck := shuffleDeck(deck)
	hands := dealCards(shuffledDeck, 4, 13)

	for i, hand := range hands {
		fmt.Printf("Hand %d:\n", i+1)
		for _, card := range hand {
			fmt.Printf("%s of %s\n", card.Value, card.Suit)
		}
		fmt.Println()
	}
}

func generateDeck() []Card {
	var deck []Card
	for _, suit := range suits {
		for _, value := range values {
			card := Card{
				Suit:  suit,
				Value: value,
			}
			deck = append(deck, card)
		}
	}
	return deck
}

func shuffleDeck(deck []Card) []Card {
	shuffledDeck := make([]Card, len(deck))
	perm := rand.Perm(len(deck))
	for i, j := range perm {
		shuffledDeck[i] = deck[j]
	}
	return shuffledDeck
}

func dealCards(deck []Card, numPlayers, cardsPerPlayer int) [][]Card {
	hands := make([][]Card, numPlayers)
	for i := 0; i < numPlayers; i++ {
		hands[i] = make([]Card, cardsPerPlayer)
		for j := 0; j < cardsPerPlayer; j++ {
			hands[i][j] = deck[i*cardsPerPlayer+j]
		}
	}
	return hands
}
