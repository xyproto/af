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

// ReLU is the rectified linear unit.
// `x >= 0 ? x : 0`
func ReLU(x float64) float64 {
	if x >= 0 {
		return x
	}
	return 0
}

// PReLU is the parametric rectified linear unit.
// `x >= 0 ? x : a * x`
func PReLU(x, a float64) float64 {
	if x >= 0 {
		return x
	}
	return a * x
}
