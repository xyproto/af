package nn

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a, b interface{}) {
	if a == b {
		return
	}
	t.Fatal(fmt.Sprintf("%v != %v", a, b))
}

func TestNN(t *testing.T) {
	v := []float64{1.0, 0.8, 0.0, 0.1, 0.2, 0.5, 0.7}
	n := New(v, []int{10, 10, 10, 1}) // Input layer, three layers 10 big, then output layer that is 1 big
	n.Process(2)                      // look two steps back for each neuron
	assertEqual(t, n.Output(), 0.2610819459019918)
}
