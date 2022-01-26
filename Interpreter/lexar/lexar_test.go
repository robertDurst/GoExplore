package lexar

// import (
// 	"testing"
// )

// func TestEmptyInputLexar(t *testing.T) {
// 	le := CreateLexarExecutor()
// 	ls, err := le.Lex("")
// 	if err != nil {
// 		t.Errorf("did not expected an error")
// 	}

// 	if len(ls) != 0 {
// 		t.Errorf("Expected 0 lexicons from empty input. Received %d.", len(ls))
// 	}
// }

// func TestAtomSExpression(t *testing.T) {
// 	le := CreateLexarExecutor()
// 	ls, err := le.Lex("ATOM")
// 	if err != nil {
// 		t.Errorf("did not expected an error")
// 	}

// 	if len(ls) != 1 {
// 		t.Errorf("Expected 1 top-level lexicon. Received %d.", len(ls))
// 	}

// 	if ls[0].Type != Atom {
// 		t.Errorf("Expected top-level lexicon to be an Atom. Received %d", ls[0].Type)
// 	}

// 	if !ls[0].IsSExpression {
// 		t.Error("Expected an SExpression.")
// 	}
// }

// func TestAtomNotSExpression(t *testing.T) {
// 	le := CreateLexarExecutor()
// 	ls, err := le.Lex("atom")
// 	if err != nil {
// 		t.Errorf("did not expected an error")
// 	}

// 	if len(ls) != 1 {
// 		t.Errorf("Expected 1 top-level lexicon. Received %d.", len(ls[0].ListValues))
// 	}

// 	if ls[0].Type != Identifier {
// 		t.Errorf("Expected top-level lexicon to be an Atom. Received %d.", ls[0].Type)
// 	}

// 	if ls[0].IsSExpression {
// 		t.Error("Did not expect an SExpression.")
// 	}
// }

// func TestAtomPanicOnMixedCasing(t *testing.T) {
// 	AssertError(t, "aTOM")
// }

// func TestListWithAtoms(t *testing.T) {
// 	le := CreateLexarExecutor()
// 	ls, err := le.Lex("(HELLO WORLD123)")
// 	if err != nil {
// 		t.Errorf("did not expected an error")
// 	}

// 	if len(ls) != 1 {
// 		t.Errorf("Expected 1 top-level lexicon. Received %d.", len(ls))
// 	}

// 	if len(ls[0].ListValues) != 2 {
// 		t.Errorf("Expected 2 atoms in list. Received %d.", len(ls[0].ListValues))
// 	}

// 	firstElement := ls[0].ListValues[0]
// 	if firstElement.Type != Atom {
// 		t.Errorf("Expected first element in list to be an Atom. Received %d.", firstElement.Type)
// 	}

// 	if firstElement.Value != "HELLO" {
// 		t.Errorf("Expected value of first element to be 'HELLO'. Received %s.", firstElement.Value)
// 	}
// }

// func TestListWithinList(t *testing.T) {
// 	le := CreateLexarExecutor()
// 	ls, err := le.Lex("(HELLO (WORLD5 ((BAR))) FOO)")
// 	if err != nil {
// 		t.Errorf("did not expected an error")
// 	}

// 	if len(ls) != 1 {
// 		t.Errorf("Expected 1 top-level lexicon. Received %d.", len(ls))
// 	}

// 	if len(ls[0].ListValues) != 3 {
// 		t.Errorf("Expected 3 atoms in list. Received %d.", len(ls[0].ListValues))
// 	}

// 	secondElement := ls[0].ListValues[1]
// 	if secondElement.Type != List {
// 		t.Errorf("Expected second element in list to be an List. Received %d.", secondElement.Type)
// 	}

// 	if len(secondElement.ListValues) != 2 {
// 		t.Errorf("Expected value of second element to be list of length 2. Received %d.", len(secondElement.ListValues))
// 	}
// }

// func TestPanicWhenArgListIsElementOfSExpressionList(t *testing.T) {
// 	AssertError(t, "(HELLO WORLD [WHAT UP])")
// }

// func TestPanicWhenMissingArgListEndParen(t *testing.T) {
// 	AssertError(t, "label[foo;label[foobar;@[[a;b;c;d;e];label[cons;bar]]]")
// }

// func TestPanicWhenMissingSExpListEndParen(t *testing.T) {
// 	AssertError(t, "((hello)")
// }

// func TestPanicWhenMissingArgListTooManyEndParen(t *testing.T) {
// 	AssertError(t, "label[foo;label[foobar;@[[a;b;c;d;e];label[cons;bar]]]]]")
// }

// func TestPanicWhenMissingSExpListTooManyEndParen(t *testing.T) {
// 	AssertError(t, "((hello)))")
// }

// func AssertError(t *testing.T, input string) {
// 	defer func() { recover() }()

// 	le := CreateLexarExecutor()
// 	_, err := le.Lex(input)

// 	if err == nil {
// 		t.Errorf("expected an error")
// 	}
// }
