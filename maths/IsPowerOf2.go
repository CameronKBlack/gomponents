package maths

func IsPowerOf2(n int) bool {
	return (n > 0) && (n&(n-1) == 0)
}