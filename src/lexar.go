package GoExplore

import (
	"bytes"
	"errors"
	"strings"
	"unicode"
)

type LexarExecutor struct {
	listStack []Lexicon
	lexs      []Lexicon
}

func (le *LexarExecutor) addNonListLexiconToLexs(lex Lexicon) {
	if len(le.listStack) == 0 {
		le.lexs = append(le.lexs, lex)
	} else {
		i := len(le.listStack) - 1
		le.listStack[i].ListValues = append(le.listStack[i].ListValues, lex)
	}
}

func (le *LexarExecutor) addTopListLexiconToLexs() error {
	length := len(le.listStack)
	if length == 0 {
		return errors.New("too many right parenthesis")
	}

	top := le.listStack[length-1]
	le.listStack = le.listStack[:length-1]

	if length > 1 &&
		top.Type == ArgList &&
		le.listStack[length-2].Type == List {
		return errors.New("argument list cannot be an element of a SExpression list")
	}

	le.addNonListLexiconToLexs(top)

	return nil
}

func (le *LexarExecutor) addListToStack(lex Lexicon) {
	le.listStack = append(le.listStack, lex)
}

func (le *LexarExecutor) Reset() {
	le.listStack = make([]Lexicon, 0)
	le.lexs = make([]Lexicon, 0)
}

func (le *LexarExecutor) Lex(line string) ([]Lexicon, error) {
	le.Reset()

	for curIndex := 0; curIndex < len(line); curIndex++ {
		curChar := line[curIndex]

		switch curChar {

		// comments (ignored)
		case '?':
			for curChar != '\n' {
				curIndex++
				if len(line) <= curIndex {
					break
				}
				curChar = line[curIndex]
			}

		// SExpression List
		case '(':
			le.addListToStack(CreateList())
		case ')':
			err := le.addTopListLexiconToLexs()
			if err != nil {
				return nil, err
			}

		// Argument List
		case '[':
			le.addListToStack(CreateArgList())
		case ']':
			err := le.addTopListLexiconToLexs()
			if err != nil {
				return nil, err
			}
		// Punctuation
		case ';':
			le.addNonListLexiconToLexs(CreateSemicolon())
		case '~':
			le.addNonListLexiconToLexs(CreateSquiggle())
		case '@':
			le.addNonListLexiconToLexs(CreateAtSign())

		// Atom/Identifier
		default:
			if unicode.IsLetter(rune(curChar)) {

				var curAtom bytes.Buffer

				curAtom.WriteByte(curChar)
				curIndex++

				for curIndex < len(line) {
					curChar = line[curIndex]
					if !isAlphaNumericByte(curChar) {
						break
					}

					curAtom.WriteByte(curChar)
					curIndex++
				}

				isSExpression := strings.ToUpper(curAtom.String()) == curAtom.String()
				if !isSExpression && (strings.ToLower(curAtom.String()) != curAtom.String()) {
					return nil, errors.New("atom must be either all upper case or all lower case")
				}

				var atom Lexicon
				if isSExpression {
					atom = CreateAtom(curAtom.String())
				} else {
					atom = CreateIdent(curAtom.String())
				}

				le.addNonListLexiconToLexs(atom)

				// back track
				curIndex--
			}
		}
	}

	if len(le.listStack) != 0 {
		return nil, errors.New("forgot to close a list")
	}

	return le.lexs, nil
}

func CreateLexarExecutor() LexarExecutor {
	return LexarExecutor{listStack: make([]Lexicon, 0), lexs: make([]Lexicon, 0)}
}

func isAlphaNumericByte(char byte) bool {
	return unicode.IsNumber(rune(char)) || unicode.IsLetter(rune(char))
}
