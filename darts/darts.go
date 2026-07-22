package darts

import "math"

func Score(x, y float64) int {
	var res int

	zone := []int{10, 5, 1, 0}
	dist := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	if dist <= 1 {
		res = zone[0]
	} else if dist <= 5 {
		res = zone[1]
	} else if dist <= 10 {
		res = zone[2]
	} else {
		res = zone[3]
	}

	return res
}
