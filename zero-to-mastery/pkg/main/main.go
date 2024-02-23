package main

import (
	"zero-to-mastery/go/pkg/display"
	"zero-to-mastery/go/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("an exiting message")
}
