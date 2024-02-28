package main

import "fmt"

type Ints32 interface {
	~int32 | ~uint32 // type approximation
}

type MyInt int32

func SumNumbers[ints Ints32](arr []ints) ints {
	var sum ints
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	nums := []int32{1, 2, 3}
	nums2 := []uint32{1, 2, 3}
	nums3 := []MyInt{MyInt(1), MyInt(2), MyInt(3)}
	fmt.Println(SumNumbers(nums))
	fmt.Println(SumNumbers(nums2))
	fmt.Println(SumNumbers(nums3))
}
