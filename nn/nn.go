package nn

import (
	"github.com/xyproto/activation/neuron"

	"fmt"
)

type NN []*neuron.Layer

// New takes a vector of inputs or a layer of inputs
func New(inputs neuron.VectorOrLayer, layersizes []int, outputsize int) *NN {
	var nn NN
	prevLayer := inputs
	for _, layersize := range layersizes {
		l := neuron.NewLayer(prevLayer, layersize)
		nn = append(nn, l)
		prevLayer = l
	}
	// Output layer, 1 largs
	l := neuron.NewLayer(prevLayer, 1)
	nn = append(nn, l)
	if layersizes[len(layersizes)-1] != 1 {
		fmt.Println("warning: last layer isn't of size 1")
	}
	return &nn
}

// Take the inputs for the neural net and generate the outputs for the neurons
// stepsBack is the number of previous outputs to add to the sum
func (n *NN) Process(stepsBack int) {
	for _, l := range *n {
		l.Process(stepsBack)
	}
}

// Get the output value from the output layer
func (n *NN) Output() float64 {
	var layers []*neuron.Layer = []*neuron.Layer(*n)
	outputLayer := layers[len(layers)-1]
	if len(*outputLayer) != 1 {
		fmt.Println("warning: output layer isn't of size 1")
	}
	return outputLayer.Outputs()[0]
}
