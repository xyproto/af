package neuron

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
