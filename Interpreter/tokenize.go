package interpreter

import (
	"GoExplore/interpreter/lexicons"
	"GoExplore/interpreter/tokens"
	"fmt"
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
		conditionalStatement := t.parseConditionalStatement(cur)
		form := tokens.CreateForm(conditionalStatement)
		t.tks = append(t.tks, form)
		t.i++
		return

	case "AtSign":
		t.i++
		functionAtSign := t.parseFunctionAtSign(t.lexs[t.i])
		form := tokens.CreateForm(functionAtSign)
		t.tks = append(t.tks, form)
		t.i++
		return
	}
}

func (t Tokenizer) parseForm2(lexs []lexicons.Lexicon) tokens.Form {
	cur := lexs[0]

	switch cur.GetType() {
	case "Atom":
		atom := cur.(lexicons.Atom)

		if atom.IsSExpression() {
			constant := tokens.CreateConstant(atom)
			return tokens.CreateForm(constant)
		} else {
			identifier := tokens.CreateIdentifier(atom.GetValue())
			if _, ok := t.knownFunctionNames[atom.GetValue()]; ok {
				// function starting identifier
				functionIdentifier := tokens.CreateFunctionIdentifier(identifier)
				return tokens.CreateForm(functionIdentifier)
			} else if atom.GetValue() == "label" {
				functionLabel := t.parseFunctionLabel(lexs[1])
				return tokens.CreateForm(functionLabel)
			} else {
				// variable
				variable := tokens.CreateVariable(identifier)
				return tokens.CreateForm(variable)
			}
		}

	case "List":
		// lists are SExpressions, thus must be a constant
		constant := tokens.CreateConstant(cur.(lexicons.List))
		return tokens.CreateForm(constant)

	case "ArgList":
		conditionalStatement := t.parseConditionalStatement(lexs[0])
		return tokens.CreateForm(conditionalStatement)

	case "AtSign":
		functionAtSign := t.parseFunctionAtSign(lexs[1])
		return tokens.CreateForm(functionAtSign)

	default:
		panic("Unexpected lexicon received in ParseForm2")
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
		panic(fmt.Sprintf("Expected Atom to be first element in 'label'. Received [%s]: %s.", identifierValue.GetType(), identifierValue))
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

// @[<var list>;<form>] |
// IN: ArgList
func (t Tokenizer) parseFunctionAtSign(cur lexicons.Lexicon) tokens.Function {
	if cur.GetType() != "ArgList" {
		panic(fmt.Sprintf("Expected ArgList to follow 'AtSign'. Received %s.", cur.GetType()))
	}

	argList := cur.(lexicons.ArgList)

	if argList.Value[0].GetType() != "ArgList" {
		panic("Expected Atom to be first element in 'label' ArgList")
	}

	varListArgList := argList.Value[0].(lexicons.ArgList)
	varList := tokens.CreateVarList()
	i := 0
	for i < len(varListArgList.Value) {
		if varListArgList.Value[i].GetType() != "Atom" || varListArgList.Value[i].(lexicons.Atom).IsSExpression() {
			panic("Expected only variables (a.k.a. identifiers) in VarList")
		}
		variable := tokens.CreateIdentifier(varListArgList.Value[i].(lexicons.Atom).GetValue())

		i++

		if i < len(varListArgList.Value) {
			if varListArgList.Value[i].GetType() != "Semicolon" {
				panic("Expected ';' to separate variables in VarList")
			}

			i++
		}
		varList.Value = append(varList.Value, variable)
	}

	if argList.Value[1].GetType() != "Semicolon" {
		panic("Expected Identifier of FunctionAtSign's VarList to be followed by a Semicolon")
	}

	form := t.parseForm2(argList.Value[2:])
	functionAtSign := tokens.CreateFunctionAtSign(varList, form)
	return functionAtSign
}

func (t Tokenizer) parseConditionalStatement(cur lexicons.Lexicon) tokens.ConditionalStatement {
	if cur.GetType() != "ArgList" {
		panic(fmt.Sprintf("Expected ArgList for ConditionalStatement. Received %s.", cur.GetType()))
	}

	conditionalStatement := tokens.CreateConditionalStatement()
	argList := cur.(lexicons.ArgList)

	i := 0
	for i < len(argList.Value) {

		predicateFormArgs := make([]lexicons.Lexicon, 0)
		for {
			if i >= len(argList.Value) {
				panic("Expected ConditionalPair of form: FORM~FORM")
			}

			if argList.Value[i].GetType() == "Squiggle" {
				i++
				break
			} else {
				predicateFormArgs = append(predicateFormArgs, argList.Value[i])
			}

			i++
		}

		resultFormArgs := make([]lexicons.Lexicon, 0)
		for {
			if i >= len(argList.Value) {
				break
			}

			if argList.Value[i].GetType() == "Semicolon" {
				i++
				break
			} else {
				resultFormArgs = append(resultFormArgs, argList.Value[i])
			}

			i++
		}

		predicate := t.parseForm2(predicateFormArgs)
		result := t.parseForm2(resultFormArgs)
		conditionalPair := tokens.CreateConditionalPair(predicate, result)
		conditionalStatement.ConditionalPairs = append(conditionalStatement.ConditionalPairs, conditionalPair)
	}

	return conditionalStatement
}
