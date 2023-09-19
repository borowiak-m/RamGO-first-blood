package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

var pl = fmt.Println
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

	if err == nil {
		pl(concat("Message received: ", message))
	} else {
		log.Fatal(err)
	}

	//break line
	pl("==================================")

	sendsSoFar, sendsToAdd := 430, 35
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	pl("Status: You've sent", sendsSoFar, "messages")

}
