package interpreter

import (
	"bytes"
	"strings"
	"unicode"
)

func lex(line string) []SExpression {
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
