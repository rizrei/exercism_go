package isbnverifier

import "strings"

const ISBNLength = 10

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != ISBNLength {
		return false
	}

	sum := 0

	for i, r := range isbn {
		if !isValidRune(r, i) {
			return false
		}

		sum += runeToInt(r) * (ISBNLength - i)
	}

	return sum%11 == 0
}

func isValidRune(r rune, index int) bool {
	return (r >= '0' && r <= '9') || (r == 'X' && index == ISBNLength-1)
}

func runeToInt(r rune) int {
	if r == 'X' {
		return 10
	}

	return int(r - '0')
}
