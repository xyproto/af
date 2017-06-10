package activation

import (
	"math"
)

func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func Tanh(x float64) float64 {
	return math.Tanh(x)
}

func ReLU(x float64) float64 {
	if x > 0 {
		return x
	}
	return 0
}

func PReLU(x, a float64) float64 {
	if x > 0 {
		return x
	}
	return x * a
}
