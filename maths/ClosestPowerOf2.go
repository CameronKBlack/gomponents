package maths

import (
	"errors"
	"fmt"
	"math/bits"
)

func ClosestPreviousPowerOf2(n uint) uint {
	return 1 << (bits.Len(n) - 1)
}

func ClosestNextPowerOf2(n uint) uint {
	return 1 << (bits.Len(n))
}

func ClosestPowerOf2(n uint) (uint, error) {
	if n == 0 {
		return 0, errors.New("input must be greater than zero")
	}
	
	next := ClosestNextPowerOf2(n)
	nextDiff := next - n
	previous := ClosestPreviousPowerOf2(n)
	previousDiff := n - previous

	if nextDiff == previousDiff {
		return previous, fmt.Errorf("input is equidistant between two powers of two (%d and %d). Lower value given", previous, next)
	}

	if nextDiff < previousDiff {
		return next, nil
	}
	return previous, nil
}