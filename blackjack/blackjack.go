package blackjack

import (
	"strings"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var res int

	switch strings.ToLower(card) {
	case "ace":
		res = 11
	case "king", "queen", "jack":
		res = 10
	case "ten":
		res = 10
	case "nine":
		res = 9
	case "eight":
		res = 8
	case "seven":
		res = 7
	case "six":
		res = 6
	case "five":
		res = 5
	case "four":
		res = 4
	case "three":
		res = 3
	case "two":
		res = 2
	case "one":
		res = 1
	default:
		res = 0
	}
	return res
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var res string
	myScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	switch {
	case card1 == card2 && card1 == "ace":
		res = "P"
	case myScore == 21 && dealerScore < 10:
		res = "W"
	case myScore == 21:
		res = "S"
	case myScore < 21 && myScore >= 17:
		res = "S"
	case myScore < 17 && myScore >= 12:
		if dealerScore >= 7 {
			res = "H"
		} else {
			res = "S"
		}
	default:
		res = "H"
	}

	return res
}
