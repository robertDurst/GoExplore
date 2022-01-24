package interpreter

import (
	"GoExplore/interpreter/errors"
	"GoExplore/interpreter/lexicons"
	"GoExplore/interpreter/tokens"
	"fmt"
)

type Tokenizer struct {
	errorHandler errors.ErrorHandler
}

func CreateTokenizer() Tokenizer {
	return Tokenizer{errorHandler: errors.CreateStdErrorHandler()}
}

func (t Tokenizer) tokenize(ls []lexicons.Lexicon) {
	parseForm(ls)
}

func parseForm(ls []lexicons.Lexicon) ([]lexicons.Lexicon, tokens.Form) {
	fmt.Printf("Parsing form...[%s]\n", ls[0].GetType())

	cur := ls[0]
	switch cur.GetType() {
	case "LParen":
		_, sexp := parseSExpression(ls)
		fmt.Println(sexp)
	case "Atom":
		if cur.GetValue() == "label" {
			parseLabeledFunction(ls[1:])
		} else {
			// TODO: check if lowercase, which then means it is a variable
			ls, _ = parseSExpression(ls)
		}
		// case "LSquareBracket":
		// 	ls, _ = parseCondition(ls[1:])
		// case "AtSign":
		// 	ls, _ = parseFunction(ls[1:])
	}

	return ls, tokens.Form{}
}

func parseSExpression(ls []lexicons.Lexicon) ([]lexicons.Lexicon, tokens.SExpression) {
	if ls[0].GetType() == "Atom" {
		return ls[1:], tokens.SExpression{Value: ls[0].GetValue()}
	} else if ls[0].GetType() == "LParen" {
		list := make([]tokens.SExpression, 0)
		ls, sexp := parseSExpression(ls[1:])
		list = append(list, sexp)
		if ls[0].GetType() == "Dot" {
			ls, sexp := parseSExpression(ls[1:])
			list = append(list, sexp)
			if ls[0].GetType() != "RParen" {
				panic("Unexpected lexicon")
			}

			lAsString := "("
			for _, l := range list {
				lAsString += fmt.Sprintf("%s ", l.Value)
			}
			lAsString = lAsString[:len(lAsString)-1]
			lAsString += ")"

			return ls[1:], tokens.SExpression{Value: lAsString}
		} else {
			for ls[0].GetType() != "RParen" {
				lss, sexp := parseSExpression(ls)

				ls = lss
				list = append(list, sexp)
			}

			lAsString := "("
			for _, l := range list {
				lAsString += fmt.Sprintf("%s ", l.Value)
			}
			lAsString = lAsString[:len(lAsString)-1]
			lAsString += ")"

			return ls[1:], tokens.SExpression{Value: lAsString}
		}
	} else {
		panic("Unexpected lexicon")
	}
}

func parseCondition(ls []lexicons.Lexicon) {
	fmt.Println("Parse parseCondition")
}

func parseFunction(ls []lexicons.Lexicon) {
	fmt.Println("Parse parseFunction")

	if ls[0].GetType() == "AtSign" {
		parseAtSignFunction(ls[1:])
	} else if ls[0].GetType() == "Atom" {
		if ls[0].GetValue() == "label" {
			parseLabeledFunction(ls[1:])
		} else {
			identifier := parseIdentifier(ls[0])
			fmt.Printf("identifier: [%s]\n", identifier)
		}
	} else {
		panic("Unexpected type, expected an AtSign or Atom")
	}
}

func parseAtSignFunction(ls []lexicons.Lexicon) {
	fmt.Println("Parse parseAtSignFunction")

	if ls[0].GetType() != "LSquareBracket" {
		panic("Expected LSquareBracket")
	}

	ls, vars := parseVarList(ls[1:])
	fmt.Printf("[%d] variables in var list\n", len(vars))

	if ls[0].GetType() != "Semicolon" {
		panic(fmt.Sprintf("Expected semicolon, received [%s]", ls[0].GetType()))
	}

	ls, form := parseForm(ls[1:])

	fmt.Printf("form: [%s]\n", form)

	if ls[0].GetType() != "RSquareBracket" {
		panic("Expected RSquareBracket")
	}
}

func parseLabeledFunction(ls []lexicons.Lexicon) {
	fmt.Println("Parse parseLabeledFunction")

	if ls[0].GetType() != "LSquareBracket" {
		panic("Expected LSquareBracket")
	}

	if ls[1].GetType() != "Atom" {
		panic("Expected Atom")
	}

	identifier := parseIdentifier(ls[1])

	fmt.Printf("identifier: [%s]\n", identifier)

	if ls[2].GetType() != "Semicolon" {
		panic("Expected Atom")
	}

	parseFunction(ls[3:])
}

func parseIdentifier(l lexicons.Lexicon) string {
	return l.GetValue()
}

func parseVarList(ls []lexicons.Lexicon) ([]lexicons.Lexicon, []tokens.Variable) {
	vars := make([]tokens.Variable, 0)

	if ls[0].GetType() != "LSquareBracket" {
		panic("Expected LSquareBracket")
	}

	lastType := ""
	lastIndex := 1

	for _, l := range ls[1:] {
		lastIndex++
		if l.GetType() == "Atom" {
			if lastType == "Atom" {
				panic("need a semicolon between atoms")
			}
			vars = append(vars, tokens.Variable{})
		} else if l.GetType() == "Semicolon" {
			if lastType == "Semicolon" {
				panic("need an atom between semicolons")
			}
		} else if l.GetType() == "RSquareBracket" {
			break
		} else {
			panic("Unexpected token encountered when parsing variable list")
		}
	}

	return ls[lastIndex:], vars
}
