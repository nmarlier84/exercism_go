package reversestring

import "slices"

func Reverse(input string) string {
	res := []rune{}
	for _, myRune := range input {
		res = slices.Insert(res, 0, myRune)
	}
	return string(res)
}
