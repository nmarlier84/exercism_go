package raindrops

import "fmt"

func Convert(number int) string {
	/*
		- is divisible by 3, add "Pling" to the result.
		- is divisible by 5, add "Plang" to the result.
		- is divisible by 7, add "Plong" to the result.
		- **is not** divisible by 3, 5, or 7, the result should be the number as a string.
	*/

	var result string // Explicitly typed
	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result = fmt.Sprintf("%d", number)
	}

	return result
}
