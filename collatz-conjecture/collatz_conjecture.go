package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var res int
	var step int

	if n <= 0 {
		return 0, errors.New("Error %d generate infinite loop")
	}

	for step = n; step != 1 || res >= 1000; res++ {
		if step%2 == 0 {
			step = step / 2
		} else {
			step = step*3 + 1
		}
	}
	if res < 1000 {
		return res, nil
	} else {
		return 0, errors.New("Error %d don't find its path to 1")
	}
}
