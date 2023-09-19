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

func concat(s1, s2 string) string {
	pl("Received var types: s1 as", typeOf(s1), "and s2 as", typeOf(s2))
	return s1 + s2
}

func incrementSends(num1, num2 int) int {
	pl("Received var types: num1 as", typeOf(num1), "and num2 as", typeOf(num2))
	return num1 + num2
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
	slice1 := make([]string, 6)
	slice2 := message[:3]
	// for loop skipping index
	for i, letter := range msgLetters {

		if i < len(slice1) {
			slice1[i] = letter
		}

		pl("Letter :", letter)
	}

	for _, rn := range msgRunes {
		pl("Rune :", rn)
	}

	for _, sl := range slice1 {
		pl("Slice char :", sl)
	}

	pl("Index 0 :", msgLetters[0])
	pl("Array length :", len(msgLetters))
	pl("Slice1 :", slice1, " of type :", typeOf(slice1))
	pl("Slice2 :", slice2, " of type :", typeOf(slice2))
	//break line
	pl("================Time==============")

	now := time.Now()

	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	pl("Random :", r1.Intn(100))

}
