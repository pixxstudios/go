package main

func Sum(nums ...int) int {
	total := 0

	for _, a := range nums {
		total = total + a
	}

	return total
}
