package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println
var pf = fmt.Printf
var typeOf = reflect.TypeOf

func main() {
	// using vars and func from another go file
	pl("Name :", name)
	intArr := []int{2, 3, 4, 5, 888}
	strArr := intArrToStrArr(intArr)
	pf("String array : %s of type of %T", strArr, strArr)
}
