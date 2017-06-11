package ltsm

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

func TestHi(t *testing.T) {
	assertEqual(t, hi(), "hi")
}
