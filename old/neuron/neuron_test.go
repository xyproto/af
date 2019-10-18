package neuron

import (
	"testing"
)

func TestSimpleNeuron(t *testing.T) {
	n := NewSimpleNeuron()
	n.Process(2)
	assertEqual(t, n.output, 1.5)
}

func TestRandomWeight(t *testing.T) {
	// Pseudo-random weights without seeding, which gives the same result at every test
	r10 := RandomWeight(10)
	r100 := RandomWeight(100)
	assertEqual(t, r10, 0.27860240964517774)
	assertEqual(t, r100, 0.032912010643698086)
}
