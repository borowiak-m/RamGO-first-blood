package main

import (
	"golang.org/x/exp/constraints"
)

type Distance int32
type Velocity float64
type Number interface {
	constraints.Float | constraints.Integer
}

func clamp[num Number](value, min, max num) num {
	if value > max {
		return max
	} else if value < min {
		return min
	} else {
		return value
	}
}

func testClampFloat32() {
	var (
		min float32 = 5.5
		max float32 = 9.9
	)
	if clamp(0, min, max) != min {
		panic("clamp failed for float32")
	}
}

func main() {
	testClampFloat32()
}
