package anagram

import (
	"slices"
	"strings"
)

func OrderString(word string) string {
	word = strings.ToLower(word)
	myRunes := []rune(word)
	slices.Sort(myRunes)
	return string(myRunes)
}

func Detect(subject string, candidates []string) []string {
	var res []string

	subject = strings.ToLower(subject)
	orderedSubject := OrderString(subject)
	for _, word := range candidates {
		if strings.ToLower(word) != subject {
			if orderedSubject == OrderString(word) {
				res = append(res, word)
			}
		}
	}

	return res
}
