package interpreter

import (
	"bytes"
	"unicode"
)

func tokenize(line string) []SExpression {

	var buffer bytes.Buffer
	var commentBuffer bytes.Buffer
	isWord := false
	isComment := false
	inList := 0
	listQueue := make([]List, 0)
	sexps := make([]SExpression, 0)

	for _, char := range line {

		// ignore comments for now
		if isComment {
			if char == '\n' {
				isComment = false
				commentBuffer.Reset()
			} else {
				commentBuffer.WriteRune(char)
				continue
			}
		}

		switch char {
		case ';':
			isComment = true
			isWord = false
		case ')':
			if inList == 0 {
				panic("tried to close a list when not open")
			} else {
				// handle possible end of an atom
				if isWord {
					isWord = false
					newAtom := Atom{name: buffer.String()}
					buffer.Reset()

					if inList > 0 {
						listQueue[inList-1].sexps = append(listQueue[inList-1].sexps, newAtom)
					} else {
						sexps = append(sexps, newAtom)
					}
				}

				inList--
				if inList == 0 {
					sexps = append(sexps, listQueue[len(listQueue)-1])
				} else {
					listQueue[len(listQueue)-2].sexps = append(listQueue[len(listQueue)-2].sexps, listQueue[len(listQueue)-1])
				}
				listQueue = listQueue[:len(listQueue)-1]
			}
		case '(':
			// handle beginning of new list
			listQueue = append(listQueue, List{sexps: make([]SExpression, 0)})
			inList++

			// handle possible end of an atom
			if isWord {
				isWord = false
				newAtom := Atom{name: buffer.String()}
				buffer.Reset()

				if inList > 0 {
					listQueue[inList-1].sexps = append(listQueue[inList-1].sexps, newAtom)
				} else {
					sexps = append(sexps, newAtom)
				}
			}
		default:
			if unicode.IsNumber(char) || unicode.IsLetter(char) {
				buffer.WriteRune(char)
				isWord = true
			} else {
				if isWord {
					isWord = false
					newAtom := Atom{name: buffer.String()}
					buffer.Reset()

					if inList > 0 {
						listQueue[inList-1].sexps = append(listQueue[inList-1].sexps, newAtom)
					} else {
						sexps = append(sexps, newAtom)
					}
				}
			}
		}
	}

	if inList != 0 {
		panic("expected a closing paren")
	}

	return sexps
}
