package main

import (
	"fmt"
	"unicode"
)

type LineCallback func(line string)

func lineIterator(lines []string, callback LineCallback) {
	for i := 0; i < len(lines); i++ {
		callback(lines[i])
	}
}

func main() {
	lines := []string{
		"there are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	var letters, numbers, punctuation, spaces byte
	lineFunc := func(line string) {
		for _, el := range line {
			if unicode.IsLetter(el) {
				letters++
			}
			if unicode.IsDigit(el) {
				numbers++
			}
			if unicode.IsPunct(el) {
				punctuation++
			}
			if unicode.IsSpace(el) {
				spaces++
			}
		}
	}

	lineIterator(lines, lineFunc)
	fmt.Printf("Values of variables: letters:%v / digits:%v / puncts:%v / spaces:%v\n", letters, numbers, punctuation, spaces)
	fmt.Println(letters, "letters")
	fmt.Println(numbers, "digits")
	fmt.Println(punctuation, "punctuation marks")
	fmt.Println(spaces, "spaces")
}
