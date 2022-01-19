package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	for _, word := range words {
		parseAtom(word)
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
