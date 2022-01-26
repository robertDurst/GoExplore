package tokenizer

import (
	"GoExplore/interpreter/lexar"
	"fmt"
)

type Tokenizer struct {
	lexs               []lexar.Lexicon
	tks                []Token
	knownFunctionNames map[string]bool
	i                  int
}

func CreateTokenizer(lexs []lexar.Lexicon) Tokenizer {
	knownFunctionNames := map[string]bool{
		"cons": true,
		"cdr":  true,
		"car":  true,
		"eq":   true,
		"atom": true,
	}

	return Tokenizer{lexs: lexs, tks: make([]Token, 0), knownFunctionNames: knownFunctionNames, i: 0}
}

func (t Tokenizer) GetTokens() []Token {
	return t.tks
}

func (t Tokenizer) tokenize() []Token {

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

	switch cur.Type {
	case lexar.Atom:
		constant := CreateConstant(cur.Value)
		form := CreateForm(constant)
		t.tks = append(t.tks, form)
		t.i++
		return

	case lexar.Identifier:
		identifier := CreateIdentifier(cur.Value)
		if _, ok := t.knownFunctionNames[cur.Value]; ok {
			// function starting identifier
			functionIdentifier := CreateFunctionIdentifier(identifier)
			form := CreateForm(functionIdentifier)
			t.tks = append(t.tks, form)
		} else if cur.Value == "label" {
			t.i++
			functionLabel := t.parseFunctionLabel(t.lexs[t.i])
			form := CreateForm(functionLabel)
			t.tks = append(t.tks, form)
		} else {
			// variable
			variable := CreateVariable(identifier)
			form := CreateForm(variable)
			t.tks = append(t.tks, form)
		}

		t.i++
		return

	case lexar.List:
		// lists are SExpressions, thus must be a constant
		constant := CreateConstant(cur.Value)
		form := CreateForm(constant)
		t.tks = append(t.tks, form)
		t.i++
		return

	case lexar.ArgList:
		conditionalStatement := t.parseConditionalStatement(cur)
		form := CreateForm(conditionalStatement)
		t.tks = append(t.tks, form)
		t.i++
		return

	case lexar.AtSign:
		t.i++
		functionAtSign := t.parseFunctionAtSign(t.lexs[t.i])
		form := CreateForm(functionAtSign)
		t.tks = append(t.tks, form)
		t.i++
		return
	}
}

func (t Tokenizer) parseForm2(lexs []lexar.Lexicon) Form {
	cur := lexs[0]

	switch cur.Type {
	case lexar.Atom:
		constant := CreateConstant(cur.Value)
		return CreateForm(constant)

	case lexar.Identifier:
		identifier := CreateIdentifier(cur.Value)
		if _, ok := t.knownFunctionNames[cur.Value]; ok {
			// function starting identifier
			functionIdentifier := CreateFunctionIdentifier(identifier)
			return CreateForm(functionIdentifier)
		} else if cur.Value == "label" {
			functionLabel := t.parseFunctionLabel(lexs[1])
			return CreateForm(functionLabel)
		} else {
			// variable
			variable := CreateVariable(identifier)
			return CreateForm(variable)
		}

	case lexar.List:
		// lists are SExpressions, thus must be a constant
		constant := CreateConstant(cur.Value)
		return CreateForm(constant)

	case lexar.ArgList:
		conditionalStatement := t.parseConditionalStatement(lexs[0])
		return CreateForm(conditionalStatement)

	case lexar.AtSign:
		functionAtSign := t.parseFunctionAtSign(lexs[1])
		return CreateForm(functionAtSign)

	default:
		panic("Unexpected lexicon received in ParseForm2")
	}
}

// label[<identifier>;<function>]
// IN: ArgList
func (t Tokenizer) parseFunctionLabel(cur lexar.Lexicon) Token {
	if cur.Type != lexar.ArgList {
		panic("Expected ArgList to follow 'label'")
	}

	identifierValue := cur.ListValues[0]

	if identifierValue.Type != lexar.Identifier {
		panic(fmt.Sprintf("Expected Identifier to be first element in 'label'. Received %d.", identifierValue.Type))
	}

	identifier := CreateIdentifier(cur.Value)

	if cur.ListValues[1].Type != lexar.Semicolon {
		panic("Expected Identifier of 'label' ArgList to be followed by a Semicolon")
	}

	function := t.parseFunction(cur.ListValues[2:])

	functionLabel := CreateFunctionLabel(identifier, function)
	return functionLabel
}

/*
	Function           = FunctionIdentifier | FunctionLabel | FunctionAtSign
	FunctionIdentifier = <identifier>
	FunctionLabel      = label[<identifier>;<function>]
	FunctionAtSign     = @[<var list>;<form>] |
*/
func (t Tokenizer) parseFunction(fnLexs []lexar.Lexicon) Function {
	cur := fnLexs[0]
	if cur.Type == lexar.Identifier {
		identifier := CreateIdentifier(cur.Value)
		if _, ok := t.knownFunctionNames[cur.Value]; ok {
			// function starting identifier, thus should end here
			if len(fnLexs) > 1 {
				panic("Expected Identifier to be final lexicon for FunctionIdentifier")
			}
			return CreateFunctionIdentifier(identifier)
		} else if cur.Value == "label" {
			if len(fnLexs) != 2 && fnLexs[1].Type != lexar.ArgList {
				panic("Expected ArgList to be final lexicon for FunctionLabel")
			}
			return t.parseFunctionLabel(fnLexs[1])
		} else {
			// function starting identifier, thus should end here
			if len(fnLexs) > 1 {
				panic("Expected Identifier to be final lexicon for FunctionIdentifier")
			}
			return CreateFunctionIdentifier(identifier)
		}
	} else if fnLexs[0].Type == lexar.AtSign {
		if len(fnLexs) != 2 && fnLexs[1].Type != lexar.ArgList {
			panic("Expected ArgList to be final lexicon for FunctionAtSign")
		}
		return t.parseFunctionAtSign(fnLexs[1])
	} else {
		panic("Unexpected lexicon to start a function")
	}
}

// @[<var list>;<form>] |
// IN: ArgList
func (t Tokenizer) parseFunctionAtSign(cur lexar.Lexicon) Function {
	if cur.Type != lexar.ArgList {
		panic(fmt.Sprintf("Expected ArgList to follow 'AtSign'. Received %d.", cur.Type))
	}

	if cur.ListValues[0].Type != lexar.ArgList {
		panic("Expected Atom to be first element in 'label' ArgList")
	}

	varListArgList := cur.ListValues[0]
	varList := make([]Variable, 0)
	i := 0
	for i < len(varListArgList.Value) {
		if varListArgList.ListValues[i].Type != lexar.Atom || varListArgList.ListValues[i].IsSExpression {
			panic("Expected only variables (a.k.a. identifiers) in VarList")
		}
		variable := CreateVariable(CreateIdentifier(varListArgList.ListValues[i].Value))

		i++

		if i < len(varListArgList.Value) {
			if varListArgList.ListValues[i].Type != lexar.Semicolon {
				panic("Expected ';' to separate variables in VarList")
			}

			i++
		}
		varList = append(varList, variable)
	}

	if cur.ListValues[1].Type != lexar.Semicolon {
		panic("Expected Identifier of FunctionAtSign's VarList to be followed by a Semicolon")
	}

	form := t.parseForm2(cur.ListValues[2:])
	functionAtSign := CreateFunctionAtSign(varList, form)
	return functionAtSign
}

func (t Tokenizer) parseConditionalStatement(cur lexar.Lexicon) ConditionalStatement {
	if cur.Type != lexar.ArgList {
		panic(fmt.Sprintf("Expected ArgList for ConditionalStatement. Received %d.", cur.Type))
	}

	conditionalStatement := CreateConditionalStatement()

	i := 0
	for i < len(cur.ListValues) {

		predicateFormArgs := make([]lexar.Lexicon, 0)
		for {
			if i >= len(cur.ListValues) {
				panic("Expected ConditionalPair of form: FORM~FORM")
			}

			if cur.ListValues[i].Type == lexar.Squiggle {
				i++
				break
			} else {
				predicateFormArgs = append(predicateFormArgs, cur.ListValues[i])
			}

			i++
		}

		resultFormArgs := make([]lexar.Lexicon, 0)
		for {
			if i >= len(cur.ListValues) {
				break
			}

			if cur.ListValues[i].Type == lexar.Semicolon {
				i++
				break
			} else {
				resultFormArgs = append(resultFormArgs, cur.ListValues[i])
			}

			i++
		}

		predicate := t.parseForm2(predicateFormArgs)
		result := t.parseForm2(resultFormArgs)
		conditionalPair := CreateConditionalPair(predicate, result)
		conditionalStatement.ConditionalPairs = append(conditionalStatement.ConditionalPairs, conditionalPair)
	}

	return conditionalStatement
}
