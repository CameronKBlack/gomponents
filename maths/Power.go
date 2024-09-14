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