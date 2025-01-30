package mathutil

import (
	"errors"
	"math"
)

func Factorial(a int) int {
	if a == 0 {
		return 1
	} else if a == 1 {
		return 1
	} else {
		return a * Factorial(a-1)
	}
}

func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("cannot find square root of negative number")
	}
	return math.Sqrt(num), nil
}
