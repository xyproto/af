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
)

func Linear(x float64) float64 { return x }
func Inv(x float64) float64    { return -x }
func Sin(x float64) float64    { return math.Sin(math.Pi * x) }
func Cosine(x float64) float64 { return math.Cos(math.Pi * x) }

func ReLu(x float64) float64 {
	if x < 0 {
		return 0
	} else {
		return x
	}
}

func UnsignedStep(x float64) float64 {
	if x > 0 {
		return 1
	} else {
		return 0
	}
}

// PReLU is the parametric rectified linear unit.
// `x >= 0 ? x : a * x`
func PReLU(x, a float64) float64 {
	if x >= 0 {
		return x
	}
	return a * x
}
