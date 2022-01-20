package interpreter

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func tokenize(line string) []SExpression {
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

func interpretLine(line string) []SExpression {
	tokens := tokenize(line)

	sexps := make([]SExpression, 0)
	inList := 0
	listQueue := make([]List, 0)
	for _, token := range tokens {

		switch token.GetName() {
		case "Atom":
			if inList > 0 {
				listQueue[inList-1].sexps = append(listQueue[inList-1].sexps, token)
			} else {
				sexps = append(sexps, token)
			}
		case "LParen":
			listQueue = append(listQueue, List{sexps: make([]SExpression, 0)})
			inList++
		case "RParen":
			if inList == 0 {
				panic("tried to close a list when not open")
			} else {
				inList--
				if inList == 0 {
					sexps = append(sexps, listQueue[len(listQueue)-1])
				} else {
					listQueue[len(listQueue)-2].sexps = append(listQueue[len(listQueue)-2].sexps, listQueue[len(listQueue)-1])
				}
				listQueue = listQueue[:len(listQueue)-1]
			}
		}
	}

	if inList != 0 {
		panic("expected a closing paren")
	}

	return sexps
}

func getNextInput() string {
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return text
}

func RunGoExplore() {
	fmt.Println("Starting GoExplore...")

	for {
		line := getNextInput()
		sexps := interpretLine(line)

		for _, sexp := range sexps {
			fmt.Println(sexp)
		}
	}
}
