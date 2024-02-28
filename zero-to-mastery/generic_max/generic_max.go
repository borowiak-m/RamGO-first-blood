package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type NumsArr[ordType constraints.Ordered] struct {
	inner []ordType
}

func (arr *NumsArr[ordType]) Max() ordType {
	max := arr.inner[0]
	for i := 1; i < len(arr.inner); i++ {
		if arr.inner[i] > max {
			max = arr.inner[i]
		}
	}
	return max
}

func main() {
	arr := NumsArr[int]{inner: []int{6, 4, 8, 9, 4, 0}}
	fmt.Println(arr.Max())
}
