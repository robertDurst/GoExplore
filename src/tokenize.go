package GoExplore

import (
	"errors"
	"fmt"
)

func Tokenize(lexs []Lexicon) (Token, error) {
	if len(lexs) == 0 {
		return nil, errors.New("received no lexicons to tokenize")
	}

	if lexs[0].Type == AtSign {
		functionAtSign, err := parseFunctionAtSign(lexs[1])
		if err != nil {
			return nil, err
		}

		return CreateForm(functionAtSign), nil
	}

	if lexs[0].Type == Ident && lexs[0].Value == "label" {
		functionLabel, err := parseFunctionLabel(lexs[1])
		if err != nil {
			return nil, err
		}

		return CreateForm(functionLabel), nil
	}

	form, err := parseForm(lexs[0])
	if err != nil {
		return nil, err
	}

	return form, nil
}

func parseForm(cur Lexicon) (Token, error) {
	switch cur.Type {
	case Atom:
		constant := CreateConstant(cur.Value)
		return CreateForm(constant), nil

	case Ident:
		variable := CreateVariable(cur.Value)
		return CreateForm(variable), nil

	case List:
		constant := CreateConstant(cur.Value)
		return CreateForm(constant), nil

	case ArgList:
		conditionalStatement, err := parseConditionalStatement(cur)
		return CreateForm(conditionalStatement), err

	default:
		return nil, fmt.Errorf("unexpected lexicon received in parseForm. Received %d", cur.Type)
	}
}

func parseFunctionLabel(cur Lexicon) (Token, error) {
	if err := AssertType(cur, ArgList); err != nil {
		return nil, err
	}

	identifier := CreateVariable(cur.ListValues[0].Value)

	form, err := Tokenize(cur.ListValues[1:])
	if err != nil {
		return nil, err
	}

	functionLabel := CreateFunctionLabel(identifier, form)
	return functionLabel, nil
}

func parseFunctionAtSign(cur Lexicon) (Token, error) {
	if err := AssertType(cur, ArgList); err != nil {
		return nil, err
	}

	if err := AssertType(cur.ListValues[0], ArgList); err != nil {
		return nil, err
	}

	varList := make([]Token, 0)
	for _, varArg := range cur.ListValues[0].ListValues {
		AssertType(varArg, Ident)
		variable := CreateVariable(varArg.Value)
		varList = append(varList, variable)
	}

	form, err := Tokenize(cur.ListValues[1:])
	if err != nil {
		return nil, err
	}

	functionAtSign := CreateFunctionAtSign(varList, form)
	return functionAtSign, nil
}

func parseConditionalStatement(cur Lexicon) (Token, error) {
	if err := AssertType(cur, ArgList); err != nil {
		return nil, err
	}

	conditionalStatement := CreateConditionalStatement()

	for _, curArgList := range cur.ListValues {
		AssertType(curArgList, ArgList)

		squiggleIndex := 0
		for ; squiggleIndex < len(curArgList.ListValues); squiggleIndex++ {
			if curArgList.ListValues[squiggleIndex].Type == Squiggle {
				break
			}
		}

		// TODO: test we don't have >1 squiggle
		if squiggleIndex == 0 || squiggleIndex == len(curArgList.ListValues) {
			return nil, errors.New("incorrect form for ConditionalPair")
		}

		predicate, err := Tokenize(curArgList.ListValues[0:squiggleIndex])
		if err != nil {
			return nil, err
		}

		result, err := Tokenize(curArgList.ListValues[squiggleIndex+1:])
		if err != nil {
			return nil, err
		}

		conditionalPair := CreateConditionalPair(predicate, result)
		conditionalStatement.ConditionalPairs = append(conditionalStatement.ConditionalPairs, conditionalPair)
	}

	return conditionalStatement, nil
}

func AssertType(actual Lexicon, expected int) error {
	if actual.Type != expected {
		return fmt.Errorf("[Type Mismatch]: expected %d. Received %d", expected, actual.Type)
	}

	return nil
}
