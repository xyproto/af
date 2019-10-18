package activation

import (
	"fmt"
	"math"
	"testing"
)

func assertEqual(t *testing.T, a, b interface{}) {
	if a == b {
		return
	}
	t.Fatal(fmt.Sprintf("%v != %v", a, b))
}

func round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}

func TestSigmoid(t *testing.T) {
	assertEqual(t, Sigmoid(0.0), 0.5)
}

func TestTanh(t *testing.T) {
	assertEqual(t, Tanh(0.0), 0.0)
}

func TestSoftplus(t *testing.T) {
	assertEqual(t, Softplus(0.0), 0.6931471805599453)
}
