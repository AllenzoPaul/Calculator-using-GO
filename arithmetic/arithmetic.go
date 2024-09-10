package arithmetic

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Sub(a, b float64) float64 {
	return a - b
}

func Mul(a, b float64) float64 {
	return a * b
}

func Div(a, b float64) (float64, error) {
	var err error
	if b == 0 {
		err = errors.New("cannot divide by zero")
		return 0, err
	}

	return a / b, err
}

func Mod(a, b float64) float64 {
	return math.Mod(a, b)
}
