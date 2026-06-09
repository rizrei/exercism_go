package blackjack

var cardValues = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	if value, ok := cardValues[card]; ok {
		return value
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	dealerScore := ParseCard(dealerCard)
	playerScore := ParseCard(card1) + ParseCard(card2)

	switch {
	case playerScore == 22:
		return "P"
	case playerScore == 21 && dealerScore < 10:
		return "W"
	case playerScore == 21 && dealerScore >= 10:
		return "S"
	case playerScore >= 17 && playerScore <= 20:
		return "S"
	case playerScore >= 12 && playerScore <= 16 && dealerScore < 7:
		return "S"
	default:
		return "H"
	}
}
