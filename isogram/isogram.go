package isogram

import (
	"slices"
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	myRunes := []rune{}
	for _, myRune := range word {
		if myRune != '-' && myRune != ' ' {
			if slices.Contains(myRunes, myRune) {
				return false
			}
			myRunes = append(myRunes, myRune)
		}
	}
	return true
}
