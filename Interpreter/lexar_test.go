package interpreter

import (
	"GoExplore/interpreter/lexicons"
	"testing"
)

func TestEmptyInputLexar(t *testing.T) {
	lexicons := lex("")
	if len(lexicons) != 0 {
		t.Errorf("Expected 0 lexicons from empty input. Received %d.", len(lexicons))
	}
}

func TestAtomSExpression(t *testing.T) {
	ls := lex("ATOM")

	if len(ls) != 1 {
		t.Errorf("Expected 1 top-level lexicon. Received %d.", len(ls))
	}

	if ls[0].GetType() != "Atom" {
		t.Errorf("Expected top-level lexicon to be an Atom. Received %s.", ls[0].GetType())
	}

	atom := ls[0].(lexicons.Atom)
	if !atom.IsSExpression() {
		t.Error("Expected an SExpression.")
	}
}

func TestAtomNotSExpression(t *testing.T) {
	ls := lex("atom")

	if len(ls) != 1 {
		t.Errorf("Expected 1 top-level lexicon. Received %d.", len(ls))
	}

	if ls[0].GetType() != "Atom" {
		t.Errorf("Expected top-level lexicon to be an Atom. Received %s.", ls[0].GetType())
	}

	atom := ls[0].(lexicons.Atom)
	if atom.IsSExpression() {
		t.Error("Did not expect an SExpression.")
	}
}

func TestAtomPanicOnMixedCasing(t *testing.T) {
	defer func() { recover() }()

	lex("aTOM")

	t.Errorf("Expected a panic, but did not panic.")
}

func TestListWithAtoms(t *testing.T) {
	ls := lex("(HELLO WORLD123)")

	if len(ls) != 1 {
		t.Errorf("Expected 1 top-level lexicon. Received %d.", len(ls))
	}

	list := (ls[0]).(lexicons.List)
	if len(list.GetValue()) != 2 {
		t.Errorf("Expected 2 atoms in list. Received %d.", len(list.Value))
	}

	firstElement := list.Value[0]
	if firstElement.GetType() != "Atom" {
		t.Errorf("Expected first element in list to be an Atom. Received %s.", firstElement.GetType())
	}

	atom := firstElement.(lexicons.Atom)
	if atom.GetValue() != "HELLO" {
		t.Errorf("Expected value of first element to be 'HELLO'. Received %s.", atom.GetValue())
	}
}

func TestListWithinList(t *testing.T) {
	ls := lex("(HELLO (WORLD5 ((BAR))) FOO)")

	if len(ls) != 1 {
		t.Errorf("Expected 1 top-level lexicon. Received %d.", len(ls))
	}

	list := (ls[0]).(lexicons.List)
	if len(list.GetValue()) != 3 {
		t.Errorf("Expected 3 atoms in list. Received %d.", len(list.Value))
	}

	secondElement := list.Value[1]
	if secondElement.GetType() != "List" {
		t.Errorf("Expected second element in list to be an List. Received %s.", secondElement.GetType())
	}

	list2 := secondElement.(lexicons.List)
	if len(list2.GetValue()) != 2 {
		t.Errorf("Expected value of second element to be list of length 2. Received %d.", len(list2.GetValue()))
	}
}

func TestPanicWhenArgListIsElementOfSExpressionList(t *testing.T) {
	defer func() { recover() }()

	lex("(HELLO WORLD [WHAT UP])")

	t.Errorf("Expected a panic, but did not panic.")
}
