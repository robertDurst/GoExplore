package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
	The most elementary type of S-Expression is the atomic symbol.

	Definition: An atomic symbol is a string of no more than thirty numerals and capital
	letters; the first character must be a letter.

	Examples:
		A
		APPLE
		PART2
		EXTRALONGSTRINGOFLETTERS
		A4B66XYZ2
*/
func parseAtom(word string) {
	fmt.Println(word)
}

func parse(line string) {
	words := strings.Split(line, " ")

	// split by word
	for _, word := range words {
		isWord := false

		wordAsRune := []rune(word)
		if unicode.IsLetter(wordAsRune[0]) {
			isWord = true
		}

		var buffer bytes.Buffer
		if isWord {
			buffer.WriteRune(wordAsRune[0])
		}

		// split by character of word
		for i, char := range strings.Split(word, "") {
			if isWord && i == 0 {
				continue
			}

			if char == ")" {
				break
			} else if char == "(" {
				break
			} else if char == "." {
				break
			} else if unicode.IsNumber(wordAsRune[i]) || unicode.IsLetter(wordAsRune[i]) {
				buffer.WriteRune(wordAsRune[i])
			}
		}

		fmt.Println(buffer.String())
	}
}

func interpretLine(line string) {
	parse(line)
}

func getNextInput() string {
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return text
}

func runGoExplore() {
	fmt.Println("Starting GoExplore...")

	for true {
		line := getNextInput()
		fmt.Printf("Input: %s", line)
		interpretLine(line)
	}
}

func main() {
	runGoExplore()
}
