package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type SExpression interface {
	GetName()
}

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
type Atom struct {
	name string
}

func (a Atom) GetName() {
	fmt.Printf("Atom: %s\n", a.name)
}

type Dot struct{}

func (d Dot) GetName() {
	fmt.Println("Dot")
}

type LParen struct{}

func (l LParen) GetName() {
	fmt.Println("Left Paren")
}

type RParen struct{}

func (r RParen) GetName() {
	fmt.Println("Right Paren")
}

func parse(line string) []SExpression {
	tokens := make([]SExpression, 0)

	words := strings.Split(line, " ")

	// split by word
	for _, word := range words {
		var buffer bytes.Buffer
		isWord := false
		wordAsRune := []rune(word)

		// split by character of word
		for i, char := range strings.Split(word, "") {
			if !isWord && unicode.IsLetter(wordAsRune[i]) {
				isWord = true
				buffer.WriteRune(wordAsRune[i])
				continue
			}

			if char == ")" {
				if isWord {
					isWord = false
					tokens = append(tokens, Atom{name: buffer.String()})
					buffer.Reset()
				}
				tokens = append(tokens, RParen{})
			} else if char == "(" {
				if isWord {
					isWord = false
					tokens = append(tokens, Atom{name: buffer.String()})
					buffer.Reset()
				}
				tokens = append(tokens, LParen{})
			} else if char == "." {
				if isWord {
					isWord = false
					tokens = append(tokens, Atom{name: buffer.String()})
					buffer.Reset()
				}
				tokens = append(tokens, Dot{})
			} else if unicode.IsNumber(wordAsRune[i]) || unicode.IsLetter(wordAsRune[i]) {
				buffer.WriteRune(wordAsRune[i])
			}
		}

		// in case we;re still in a word after finishing iteration here
		if isWord {
			tokens = append(tokens, Atom{name: buffer.String()})
		}
	}

	return tokens
}

func interpretLine(line string) {
	tokens := parse(line)

	for _, token := range tokens {
		token.GetName()
	}
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
