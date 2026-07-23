// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	res := []rune{}
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(strings.ReplaceAll(s, "-", " "), "_", " ")
	words := strings.Split(s, " ")

	for _, word := range words {
		runes := []rune(word)

		if len(runes) > 0 {
			res = append(res, runes[0])
		}
	}

	return string(res)
}
