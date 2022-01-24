package interpreter

import (
	"bytes"
	"unicode"
	"GoExplore/interpreter/lexicons"
)

func lex(line string) []lexicons.Lexicon {
	ls := make([]lexicons.Lexicon, 0)

	var curAtom bytes.Buffer
	lenLine := len(line)
	curIndex := 0

	for curIndex < lenLine {
		curChar := line[curIndex]

		switch curChar {

		// ignore comments
		case '?':
			for curChar != '\n' {
				curIndex++
				if (lenLine <= curIndex) {
					break
				}
				curChar = line[curIndex]
			}
		case '(':
			ls = append(ls, lexicons.LParen{})
		case ')':
			ls = append(ls, lexicons.RParen{})
		case '[':
			ls = append(ls, lexicons.LSquareBracket{})
		case ']':
			ls = append(ls, lexicons.RSquareBracket{})
		case '.':
			ls = append(ls, lexicons.Dot{})
		case ';':
			ls = append(ls, lexicons.Semicolon{})
		case '~':
			ls = append(ls, lexicons.Squiggle{})
		case '@':
			ls = append(ls, lexicons.AtSign{})
		default:
			if unicode.IsLetter(rune(curChar)) {
				curAtom.Reset()
				curAtom.WriteByte(curChar)

				curIndex++
				for curIndex <  lenLine {
					curChar = line[curIndex]
					if !isAlphaNumericByte(curChar) {
						break
					}

					curAtom.WriteByte(curChar)
					
					curIndex++
				}
				ls = append(ls, lexicons.CreateAtom(curAtom.String()))
				curIndex--
			} 
		}

		curIndex++
	}

	return ls
}

func isAlphaNumericByte(char byte) bool {
	return unicode.IsNumber(rune(char)) || unicode.IsLetter(rune(char))
}
