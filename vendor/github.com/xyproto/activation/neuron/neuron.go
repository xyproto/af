package neuron

import (
	"errors"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/xyproto/activation"
)

type (
	ActivationFunction func(float64) float64

	Layer         []*Neuron
	Vector        []float64
	VectorOrLayer interface{}

	Neuron struct {
		f              ActivationFunction
		weight         float64
		inputs         VectorOrLayer
		output         float64
		prevOutput     float64
		prevPrevOutput float64
	}
)

var (
	defaultActivationFunction = activation.ReLU
	errCalc                   = errors.New("Unrecognized inputs")
)

// Initialize the pseudo-random number generator
func Seed() {
	rand.Seed(time.Now().Unix())
}

// Create a random weight based on the number of neurons in a layer
func RandomWeight(n int) float64 {
	m := (1.0 / math.Sqrt(float64(n)))
	return (rand.Float64() * m * 2.0) - m
}

// Create a new Neuron with the default activation function, a weight of 1.0 and
// inputs of 1.0, 0.5 and 0.0.
func NewSimpleNeuron() *Neuron {
	return &Neuron{
		f:      defaultActivationFunction,
		weight: 1.0,
		inputs: Vector{1.0, 0.5, 0.0},
		output: 0.0,
	}
}

// Length of a Vector
func (v *Vector) Len() int {
	return len(*v)
}

// Length of a Layer
func (l *Layer) Len() int {
	return len(*l)
}

func AnythingToVectorOrLayer(anything interface{}) VectorOrLayer {
	if floats, ok := anything.([]float64); ok {
		return Vector(floats)
	} else if layer, ok := anything.(Layer); ok {
		return Layer(layer)
	} else if layer, ok := anything.(*Layer); ok {
		return Layer(*layer)
	} else if vector, ok := anything.(Vector); ok {
		return Vector(vector)
	} else if vector, ok := anything.(*Vector); ok {
		return Vector(*vector)
	}
	log.Fatalf("AnythingToVectorOrLayer: unrecognized input, neither []float64, Vector or Layer: %T\n", anything)
	return 0
}

func Length(v VectorOrLayer) int {
	if floats, ok := v.([]float64); ok {
		return len(floats)
	} else if layer, ok := v.(Layer); ok {
		return layer.Len()
	} else if layer, ok := v.(*Layer); ok {
		return layer.Len()
	} else if vector, ok := v.(Vector); ok {
		return vector.Len()
	} else if vector, ok := v.(*Vector); ok {
		return vector.Len()
	}
	// Should never happen
	log.Fatalf("Neuron.Length: unrecognized inputs, neither Vector nor Layer: %T\n", v)
	return 0
}

// Create a new Neuron with the given inputs, the default activation function
// and a random weight based on the size of the input vector or layer.
func New(inputs VectorOrLayer) *Neuron {
	if floats, ok := inputs.([]float64); ok {
		inputs = Vector(floats)
	} else if _, ok := inputs.(Layer); ok {
	} else if _, ok := inputs.(*Layer); ok {
	} else if _, ok := inputs.(Vector); ok {
	} else if _, ok := inputs.(*Vector); ok {
	} else {
		log.Fatalf("Neuron.New unrecognized inputs, neither Vector nor Layer: %T\n", inputs)
	}
	return &Neuron{
		f:      defaultActivationFunction,
		weight: RandomWeight(Length(inputs)),
		inputs: inputs,
		output: 0.0,
	}
}

// Process the inputs and create an output.
// Calculate a neuron output based on the neuron inputs.
// Meant to be used concurrently without having to return anything.
// stepsBack is how many previous outputs to include in the sum before
// passing it through the activation function.
// stepsBack = 1 adds the previous output
// stepsBack = 2 adds the previous output and the one before that
// stepsBack > 2 is not supported
func (n *Neuron) Process(stepsBack int) {
	sum := 0.0
	// Create a weighted sum
	if layer, ok := n.inputs.(*Layer); ok {
		for _, n := range *layer {
			sum += n.output * n.weight
		}
	} else if vector, ok := n.inputs.(*Vector); ok {
		for _, x := range *vector {
			sum += x * n.weight
		}
	} else if layer, ok := n.inputs.(Layer); ok {
		for _, n := range layer {
			sum += n.output * n.weight
		}
	} else if vector, ok := n.inputs.(Vector); ok {
		for _, x := range vector {
			sum += x * n.weight
		}
	} else if floats, ok := n.inputs.([]float64); ok {
		for _, x := range floats {
			sum += x * n.weight
		}
	} else {
		// Should never happen
		log.Fatalf("Neuron.Process: unrecognized inputs, neither Vector nor Layer: %T\n", n.inputs)
	}
	if stepsBack >= 2 {
		sum += n.prevPrevOutput
	}
	if stepsBack >= 1 {
		sum += n.prevOutput
	}
	// Save the two last outputs
	n.prevPrevOutput = n.prevOutput
	n.prevOutput = n.output
	// Pass the weighted sum through the activation function and create the output
	n.output = n.f(sum)
}

// Return the output value of the neuron.
// The output valu is calculated with the Process function.
func (n *Neuron) Output() float64 {
	return n.output
}

// New creates a new layer given several inputs and the size of the new layer
func NewLayer(inputs VectorOrLayer, size int) *Layer {
	var neurons Layer
	for i := 0; i < size; i++ {
		n := New(inputs)
		neurons = append(neurons, n)
	}
	return &neurons
}

// Take the inputs for the layer and generate the outputs for the neurons
// stepsBack is the number of previous outputs to add to the sum
func (l *Layer) Process(stepsBack int) {
	var wg sync.WaitGroup
	wg.Add(len(*l))
	for _, n := range *l {
		go func(n *Neuron) {
			defer wg.Done()
			n.Process(stepsBack)
		}(n)
	}
	wg.Wait()
}

// Get the output values from all the neurons in this layer.
// The output values are calculated with the Process function.
func (l *Layer) Outputs() []float64 {
	var outputs []float64
	for _, n := range *l {
		outputs = append(outputs, n.output)
	}
	return outputs
}
