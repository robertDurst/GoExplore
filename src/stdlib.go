package GoExplore

import (
	"errors"
)

// elementary SExpression functions
func cons(args []Token) (Token, error) {
	if len(args) != 2 && args[0].GetType() != "SExpression" && args[1].GetType() != "SExpression" {
		return CreateSExpression(CreateAtom("")), errors.New("expected only 2 arguments, each of type SExpression")
	}

	a := args[0].(SExpression)
	b := args[1].(SExpression)

	encompassingList := CreateList()
	encompassingList.ListValues = append(encompassingList.ListValues, a.Value)
	encompassingList.ListValues = append(encompassingList.ListValues, b.Value)
	return CreateSExpression(encompassingList), nil
}

func car(args []Token) (Token, error) {
	if len(args) != 1 && args[0].GetType() != "SExpression" {
		return CreateSExpression(CreateAtom("")), errors.New("expected only 1 argument of type SExpression")
	}

	a := args[0].(SExpression)

	if a.Value.Type != List || len(a.Value.ListValues) == 0 {
		return CreateSExpression(CreateAtom("")), errors.New("expected a list with at least one value")
	}

	return CreateSExpression(a.Value.ListValues[0]), nil
}

func cdr(args []Token) (Token, error) {
	if len(args) != 1 && args[0].GetType() != "SExpression" {
		return CreateSExpression(CreateAtom("")), errors.New("expected only 1 argument of type SExpression")
	}

	a := args[0].(SExpression)

	if a.Value.Type != List || len(a.Value.ListValues) == 0 {
		return CreateSExpression(CreateAtom("")), errors.New("expected a list with at least one value")
	}

	list := CreateList()
	list.ListValues = append(list.ListValues, a.Value.ListValues[1:]...)

	return CreateSExpression(list), nil
}

func eq(args []Token) (Token, error) {
	if len(args) != 2 && args[0].GetType() != "SExpression" && args[1].GetType() != "SExpression" {
		return boolToSExpression(false), errors.New("expected only 2 arguments, each of type SExpression")
	}

	a := []Token{args[0]}
	b := []Token{args[1]}

	c1, err := atom(a)
	if err != nil {
		return boolToSExpression(false), err
	}

	c2, err := atom(b)
	if err != nil {
		return boolToSExpression(false), err
	}

	if c1.(SExpression).Value.Value != "T" || c2.(SExpression).Value.Value != "T" {
		return boolToSExpression(false), errors.New("expected both arguments to be of type Atom")
	}

	return boolToSExpression(args[0].(SExpression).Value.Value == args[1].(SExpression).Value.Value), nil
}

func atom(args []Token) (Token, error) {
	if len(args) != 1 && args[0].GetType() != "SExpression" {
		return boolToSExpression(false), errors.New("expected only 1 argument of type SExpression")
	}

	a := args[0].(SExpression)

	return boolToSExpression(a.Value.Type == Atom), nil
}

func boolToSExpression(condition bool) SExpression {
	if condition {
		return CreateSExpression(CreateAtom("T"))
	}
	return CreateSExpression(CreateAtom("F"))
}
