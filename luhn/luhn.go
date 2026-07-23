package luhn

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

func Step1(id string) string {
	var res string
	digitParity := utf8.RuneCountInString(id) % 2

	for i, rune := range id {
		if i%2 == digitParity {
			tmp := int(rune-'0') * 2
			if tmp > 9 {
				tmp -= 9
			}
			res = fmt.Sprintf("%s%d", res, tmp)
		} else {
			res = fmt.Sprintf("%s%d", res, (int(rune - '0')))
		}

	}
	// fmt.Println("---------")
	// fmt.Printf("id:  %s\n", id)
	// fmt.Printf("res: %s\n", res)
	return res
}

func SumDigit(id string) int {
	var res int
	for _, rune := range id {
		res += int(rune - '0')
	}
	// fmt.Printf("sum: %d\n", res)
	// fmt.Printf("result: %t\n", res%10 == 0)
	return res
}

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	// Not enough rune in the string
	if utf8.RuneCountInString(id) <= 1 {
		return false
	}

	// Contains other rune than digit
	re := regexp.MustCompile(`[^0-9]+`)
	if re.MatchString(id) {
		return false
	}

	// Step 1: double every second char
	id = Step1(id)

	if SumDigit(id)%10 == 0 {
		return true
	}

	return false
}
