package interpreter

import (
	"bytes"
	"strings"
	"unicode"
	"GoExplore/interpreter/lexicons"
)

func lex(line string) []lexicons.Lexicon {
	ls := make([]lexicons.Lexicon, 0)

	var curAtom bytes.Buffer
	lenLine := len(line)
	curIndex := 0

	curListType := make([]int, 0) // a binary where 1 is sexp and 2 is arg
	sexpList := make([]lexicons.List, 0)
	argList := make([]lexicons.ArgList, 0)

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
			sexpList = append(sexpList, lexicons.CreateList())
			curListType = append(curListType, 1)
		case ')':
			if len(sexpList) == 0 {
				panic("Too many right parenthesis.")
			}

			curListType = curListType[:len(curListType)-1]

			if len(curListType) == 0 {
				ls = append(ls, sexpList[len(sexpList)-1])
				sexpList = sexpList[:len(sexpList)-1]
			} else if curListType[len(curListType)-1] == 2 {
				argList[len(argList)-1].Value = append(argList[len(argList)-1].Value, sexpList[len(sexpList)-1])
				sexpList = sexpList[:len(sexpList)-1]
			} else {
				sexpList[len(sexpList)-2].Value = append(sexpList[len(sexpList)-2].Value, sexpList[len(sexpList)-1])
				sexpList = sexpList[:len(sexpList)-1]
			}
			
		case '[':
			argList = append(argList, lexicons.CreateArgList())
			curListType = append(curListType, 2)
		case ']':
			if len(argList) == 0 {
				panic("Too many right square brackets.")
			}

			// TODO: fix an error which exists somewhere around here
			curListType = curListType[:len(curListType)-1]

			if len(curListType) == 0 {
				ls = append(ls, argList[len(argList)-1])
				argList = argList[:len(argList)-1]
			} else if curListType[len(curListType)-1] == 2 {
				argList[len(argList)-2].Value = append(argList[len(argList)-2].Value, argList[len(argList)-1])
				argList = argList[:len(argList)-1]
			} else {
				panic("An ArgList cannot be an element of a SExpression List.")
			}
		case ';':
			semicolon := lexicons.Semicolon{}
			if len(curListType) == 0 {
				ls = append(ls, semicolon)
			} else if curListType[len(curListType)-1] == 1 {
				sexpList[len(sexpList)-1].Value = append(sexpList[len(sexpList)-1].Value, semicolon)
			} else {
				argList[len(argList)-1].Value = append(argList[len(argList)-1].Value, semicolon)
			}		
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

				isSExpression := strings.ToUpper(curAtom.String()) == curAtom.String()
				if !isSExpression && (strings.ToLower(curAtom.String()) != curAtom.String()) {
					panic("Atom must be either all upper case or all lower case")
				}
				atom := lexicons.CreateAtom(curAtom.String(), isSExpression)

				if len(curListType) == 0 {
					ls = append(ls, atom)
				} else if curListType[len(curListType)-1] == 1 {
					sexpList[len(sexpList)-1].Value = append(sexpList[len(sexpList)-1].Value, atom)
				} else {
					argList[len(argList)-1].Value = append(argList[len(argList)-1].Value, atom)
				}
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
