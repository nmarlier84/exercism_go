package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	alphabet := map[rune]int{}

	for _, rune := range input {
		if rune >= 'a' && rune <= 'z' {
			alphabet[rune] = 1
		}
	}

	return len(alphabet) == 26
}
