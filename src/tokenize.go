package GoExplore

import (
	"errors"
	"fmt"
)

func Tokenize(lexs []Lexicon) (Token, error) {
	switch {
	case len(lexs) == 0:
		return nil, errors.New("received no lexicons to tokenize")
	case lexs[0].Type == Identifier && lexs[0].Value == "label":
		return parseFunctionLabel(lexs[1])
	default:
		if lexs[0].Type == Identifier && len(lexs) > 1 {
			name := lexs[0].Value
			args, err := parseArgs(lexs[1])
			if err != nil {
				return nil, err
			}

			return CreateFunction(name, args), nil
		}
		return parseForm(lexs[0])

	}
}

func parseForm(cur Lexicon) (Token, error) {
	switch cur.Type {
	case Atom, List:
		return CreateSExpression(cur), nil

	case Identifier:
		return CreateIdent(cur.Value), nil

	case ArgList:
		return parseConditionalStatement(cur)

	default:
		return nil, fmt.Errorf("unexpected lexicon received in parseForm. Received %d", cur.Type)
	}
}

func parseFunctionLabel(cur Lexicon) (Token, error) {
	if err := AssertType(cur, ArgList); err != nil {
		return nil, err
	}

	identifier := CreateIdent(cur.ListValues[0].Value)

	form, err := Tokenize(cur.ListValues[1:])
	if err != nil {
		return nil, err
	}

	functionLabel := CreateFunctionLabel(identifier, form)
	return functionLabel, nil
}

func parseConditionalStatement(cur Lexicon) (Token, error) {
	if err := AssertType(cur, ArgList); err != nil {
		return nil, err
	}

	conditionalStatement := CreateConditionalStatement()

	for _, curArgList := range cur.ListValues {
		if err := AssertType(curArgList, ArgList); err != nil {
			return nil, err
		}
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

func parseArgs(cur Lexicon) ([]Token, error) {
	if err := AssertType(cur, ArgList); err != nil {
		return nil, err
	}

	args := make([]Token, 0)

	for _, curArgList := range cur.ListValues {
		if err := AssertType(curArgList, ArgList); err != nil {
			return nil, err
		}

		arg, err := Tokenize(curArgList.ListValues)
		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	return args, nil
}

func AssertType(actual Lexicon, expected int) error {
	if actual.Type != expected {
		return fmt.Errorf("[Type Mismatch]: expected %d. Received %d", expected, actual.Type)
	}

	return nil
}
