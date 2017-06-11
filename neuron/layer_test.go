package neuron

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

func TestLayer(t *testing.T) {
	v := []float64{1.0, 0.8, 0.0, 0.1, 0.2, 0.5, 0.7}
	l := NewLayer(v, 1)
	l.Process(2)
	assertEqual(t, l.Outputs()[0], 0.2610819459019918)
}
