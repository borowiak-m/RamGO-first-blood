package main

import (
	"fmt"
)

func concat(s1, s2 string) string {
	return s1 + s2
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}

func main() {
	fmt.Println(concat("this is a message", " and this is its continuation"))
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")

}
