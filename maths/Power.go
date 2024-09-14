package maths

func Power(base, power int) int {
	if power == 0 {
		return 1
	}
	if power == 1 {
		return base
	}
	res := 0
	for range power - 1 {
		res += base * base
	}
	return res
}

func Sum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}