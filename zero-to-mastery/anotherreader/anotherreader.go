package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdHello   = "hello"
	CmdGoodbye = "bye"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numLines := 0
	numCommands := 0
	for scanner.Scan() {
		input := scanner.Text()
		if strings.ToUpper(input) == "Q" {
			break
		} else {
			text := strings.TrimSpace(input)
			switch text {
			case CmdHello:
				numCommands += 1
				fmt.Println("command response: hi!")
			case CmdGoodbye:
				numCommands += 1
				fmt.Println("command response: goodbye!")
			}
			if text != "" {
				numLines += 1
			}
		}
	}
	fmt.Printf("You entered %v lines \n", numLines)
	fmt.Printf("You entered %v commands \n", numCommands)
}
