package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// reader := strings.NewReader("SAMPLE")

	// var newString strings.Builder
	// buffer := make([]byte, 4)

	// for {
	// 	numBytes, err := reader.Read(buffer)
	// 	chunk := buffer[:numBytes]
	// 	newString.Write(chunk)
	// 	fmt.Printf("Read %v bytes: %c\n", numBytes, chunk)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }

	source := strings.NewReader("SAMPLE2")
	buffered := bufio.NewReader(source)
	newString, err := buffered.ReadString('\n')
	if err == io.EOF {
		fmt.Println(newString)
	} else {
		fmt.Println("something went wrong...")
	}

	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 5)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
	fmt.Printf("Line count: %v\n", len(lines))
	for _, line := range lines {
		fmt.Printf("Line: %v\n", line)
	}

}
