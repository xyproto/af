package activation

import (
	"github.com/xyproto/swish"
	"math"
)

// The swish package offers optimized Swish, Sigmoid
// and SoftPlus activation functions
var (
	Sigmoid  = swish.Sigmoid
	Swish    = swish.Swish
	SoftPlus = swish.SoftPlus
	Abs      = math.Abs
	Tanh     = math.Tanh
	Sin      = math.Sin
	Cos      = math.Cos
	Inv      = func(x float64) float64 { return -x }
)

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
