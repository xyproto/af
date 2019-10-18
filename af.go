package activation

import (
	"github.com/xyproto/swish"
	"math"
)

var Functions = [](func(float64) float64){
	swish.Sigmoid, swish.Swish, swish.SoftPlus, Tanh, ReLU,
}

var Functions2 = [](func(float64, float64) float64){
	PReLU,
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
