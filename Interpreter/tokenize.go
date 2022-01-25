package interpreter

import (
	"GoExplore/interpreter/lexicons"
	"GoExplore/interpreter/tokens"
)

type Tokenizer struct {
	lexs               []lexicons.Lexicon
	tks                []tokens.Token
	knownFunctionNames map[string]bool
	i                  int
}

func CreateTokenizer(lexs []lexicons.Lexicon) Tokenizer {
	knownFunctionNames := map[string]bool{
		"cons": true,
		"cdr":  true,
		"car":  true,
		"eq":   true,
		"atom": true,
	}

	return Tokenizer{lexs: lexs, tks: make([]tokens.Token, 0), knownFunctionNames: knownFunctionNames, i: 0}
}

func (t Tokenizer) GetTokens() []tokens.Token {
	return t.tks
}

func (t Tokenizer) tokenize() []tokens.Token {

	for t.i < len(t.lexs) {
		t.parseForm()
	}

	return t.tks
}

// <constant> |
// <variable> |
// <function>[<argument>;...;<argument>] |
// [<form>~<form>;...;<form>~<form>]
func (t *Tokenizer) parseForm() {
	cur := t.lexs[t.i]

	switch cur.GetType() {
	case "Atom":
		atom := cur.(lexicons.Atom)

		if atom.IsSExpression() {
			constant := tokens.CreateConstant(atom)
			form := tokens.CreateForm(constant)
			t.tks = append(t.tks, form)
		} else {
			identifier := tokens.CreateIdentifier(atom.GetValue())
			if _, ok := t.knownFunctionNames[atom.GetValue()]; ok {
				// function starting identifier
				functionIdentifier := tokens.CreateFunctionIdentifier(identifier)
				form := tokens.CreateForm(functionIdentifier)
				t.tks = append(t.tks, form)
			} else if atom.GetValue() == "label" {
				t.i++
				functionLabel := t.parseFunctionLabel(t.lexs[t.i])
				form := tokens.CreateForm(functionLabel)
				t.tks = append(t.tks, form)
			} else {
				// variable
				variable := tokens.CreateVariable(identifier)
				form := tokens.CreateForm(variable)
				t.tks = append(t.tks, form)
			}
		}

		t.i++
		return

	case "List":
		// lists are SExpressions, thus must be a constant
		constant := tokens.CreateConstant(cur.(lexicons.List))
		form := tokens.CreateForm(constant)
		t.tks = append(t.tks, form)
		t.i++
		return

	case "ArgList":
		panic("Not implemented")
	}
}

// label[<identifier>;<function>]
// IN: ArgList
func (t Tokenizer) parseFunctionLabel(cur lexicons.Lexicon) tokens.Token {
	if cur.GetType() != "ArgList" {
		panic("Expected ArgList to follow 'label'")
	}

	argList := cur.(lexicons.ArgList)
	identifierValue := argList.Value[0]

	if identifierValue.GetType() != "Atom" {
		panic("Expected Atom to be first element in 'label' ArgList")
	}

	atom := identifierValue.(lexicons.Atom)
	if atom.IsSExpression() {
		panic("Expected an identifier but received an SExpression")
	}

	identifier := tokens.CreateIdentifier(atom.GetValue())

	if argList.Value[1].GetType() != "Semicolon" {
		panic("Expected Identifier of 'label' ArgList to be followed by a Semicolon")
	}

	function := t.parseFunction(argList.Value[2:])

	functionLabel := tokens.CreateFunctionLabel(identifier, function)
	return functionLabel
}

// <identifier> |
// @[<var list>;<form>] |
// label[<identifier>;<function>]
// IN: ?????
func (t Tokenizer) parseFunction(fnLexs []lexicons.Lexicon) tokens.Function {

	if fnLexs[0].GetType() == "Atom" {
		atom := fnLexs[0].(lexicons.Atom)
		identifier := tokens.CreateIdentifier(atom.GetValue())
		if _, ok := t.knownFunctionNames[atom.GetValue()]; ok {
			// function starting identifier, thus should end here
			if len(fnLexs) > 1 {
				panic("Expected Identifier to be final lexicon for FunctionIdentifier")
			}
			return tokens.CreateFunctionIdentifier(identifier)
		} else if atom.GetValue() == "label" {
			if len(fnLexs) != 2 && fnLexs[1].GetType() != "ArgList" {
				panic("Expected ArgList to be final lexicon for FunctionLabel")
			}
			return t.parseFunctionLabel(fnLexs[1])
		} else {
			// function starting identifier, thus should end here
			if len(fnLexs) > 1 {
				panic("Expected Identifier to be final lexicon for FunctionIdentifier")
			}
			return tokens.CreateFunctionIdentifier(identifier)
		}
	} else if fnLexs[0].GetType() == "AtSign" {
		if len(fnLexs) != 2 && fnLexs[1].GetType() != "ArgList" {
			panic("Expected ArgList to be final lexicon for FunctionAtSign")
		}
		return t.parseFunctionAtSign(fnLexs[1])
	} else {
		panic("Unexpected lexicon to start a function")
	}
}

func (t Tokenizer) parseFunctionAtSign(fnLex lexicons.Lexicon) tokens.Function {
	return nil
}
