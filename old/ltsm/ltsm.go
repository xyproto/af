package ltsm

import (
	"github.com/xyproto/activation"
)

func hi() string {
	return "hi"
}

func BLAH(n *Neuron) {
	z := activation.Sigmoid(n * h)
	r := activation.Sigmoid(w* h)
	h := activation.Tanh(w * )
	h := (1 - z)
}