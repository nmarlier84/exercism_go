package armstrongnumbers

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	// Calculated sum
	var armSum float64

	myString := strconv.Itoa(n)
	nbDigits := float64(len(myString))

	for _, digit := range myString {
		armSum += math.Pow(float64(digit-'0'), nbDigits)
	}
	return n == int(armSum)
}
