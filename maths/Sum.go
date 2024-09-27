package maths

func Sum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}