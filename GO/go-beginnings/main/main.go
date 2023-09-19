package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
	"unicode/utf8"
)

var pl = fmt.Println
var pf = fmt.Printf
var typeOf = reflect.TypeOf

// func funcName(params) returnType {BODY}

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

// arg passed as value, copy of variable
func changeVal(myVal int) {
	myVal = 12
}

// arg passed as "ref" pointer to an var address
func changeValPoint(myPtr *int) {
	*myPtr = 12
}

func main() {
	pl("what is your message:")
	reader := bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	message = strings.TrimSpace(message)
	message = strings.Replace(message, " ", "-", -1)
	msgLetters := strings.Split(message, "")
	msgRunes := []rune(message)

	pl(concat("Message received: ", message))

	//break line
	pl("==================================")

	sendsSoFar, sendsToAdd := 430, 35
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	pl("Status: You've sent", sendsSoFar, "messages")

	//break line
	pl("===============Strings===========")

	pl("Split :", strings.Split(message, "-"))
	pl("Lower :", strings.ToLower(message))
	pl("Upper :", strings.ToUpper(message))
	pl("Rune Count :", utf8.RuneCountInString(message))

	// for each
	for i, runeVal := range message {
		pf("%d :  %#U : %c\n", i, runeVal, runeVal)
	}

	//break line
	pl("================ArrayAndSlice======")

	var slice1, slice2, slice3 []string
	var msgCharInts []int
	slice2 = strings.Split(message, "")[:3]
	slice3 = append(slice2, "ADDED")
	// for loop skipping index
	for i, letter := range msgLetters {

		if i < len(slice1) {
			slice1[i] = letter
		}

		pl("Letter :", letter)
	}

	for _, rn := range msgRunes {
		pl("Rune :", rn)
		msgCharInts = append(msgCharInts, int(rn))
	}

	for _, sl := range slice1 {
		pl("Slice char :", sl)
	}

	pl("Index 0 :", msgLetters[0])
	pl("Array length :", len(msgLetters))
	pl("Slice1 :", slice1, " of type :", typeOf(slice1))
	pl("Slice2 :", slice2, " of type :", typeOf(slice2))
	pl("Slice3 :", slice3, " of type :", typeOf(slice3))
	//break line
	pl("================Time==============")

	now := time.Now()

	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	pl("Random :", r1.Intn(100))

	//break line
	pl("================Func==============")

	// func funcName(params) returnType {BODY}
	divRes, err := getQuotient(5, 0)
	pl("Divide 5 / 0 :", divRes, "Error :", err)
	pl("Running total :", getSumOfNums(1, 2, 3, 4, 45))
	pl("Running total :", getSumOfArr(msgCharInts))
	f3 := 1
	pl("Before change :", f3)
	// run func passing by value
	changeVal(f3)
	pl("Attempt by val at changing var f3 :", f3)
	// run func passing by ref
	changeValPoint(&f3)
	pl("Attempt by val at changing var f3 :", f3)
	var f3poiner *int = &f3
	pl("var f3 address :", f3poiner)
	pl("var f3 value from pointer :", *f3poiner)
	*f3poiner = 111
	pl("var f3 value from pointer after change:", *f3poiner)
}
