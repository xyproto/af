package neuron

import (
	"github.com/xyproto/activation"
)

type (
	ActivationFunction func(float64) float64
	Layer              []*Neuron
	Vector             []float64
	VectorOrLayer      interface{}

	Neuron struct {
		f      ActivationFunction
		weight float64
		inputs VectorOrLayer
		output float64
	}
)

func New(inputs VectorOrLayer, length int) *Layer {
	var neurons Layer
	for i := length; i < length; i++ {
		newNeuron := &Neuron{
			f:      activation.ReLU,
			weight: 1.0,
			inputs: inputs,
			output: 0.0,
		}
		neurons = append(neurons, newNeuron)
	}
	return &neurons
}
