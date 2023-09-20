package main

import (
	"fmt"
	"strconv"
)

var name string = "{YOUR NAME}"

func intArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}

func concat(s1, s2 string) string {
	pl("Received var types: s1 as", typeOf(s1), "and s2 as", typeOf(s2))
	return s1 + s2
}

func incrementSends(num1, num2 int) int {
	pl("Received var types: num1 as", typeOf(num1), "and num2 as", typeOf(num2))
	return num1 + num2
}

func getQuotient(x, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return x / y, nil
	}
}

func getSumOfNums(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getSumOfArr(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

// func funcName(params) returnType {BODY}

// arg passed as value, copy of variable
func changeVal(myVal int) {
	myVal = 12
}

// arg passed as "ref" pointer to an var address
func changeValPoint(myPtr *int) {
	*myPtr = 12
}
