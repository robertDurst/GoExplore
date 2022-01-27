package GoExplore

import (
	"testing"
)

func TestCons_TwoAtoms(t *testing.T) {
	a := CreateSExpression(CreateAtom("a"))
	b := CreateSExpression(CreateAtom("b"))
	c, err := cons([]Token{a, b})
	if err != nil {
		t.Errorf("did not expect error")
	}

	if c.(SExpression).Value.Type != List {
		t.Errorf("Expected a list type. Received %d", c.(SExpression).Value.Type)
	}

	if len(c.(SExpression).Value.ListValues) != 2 {
		t.Errorf("Expected cons'd list to have length of 2. Received %d", len(c.(SExpression).Value.ListValues))
	}

	if c.(SExpression).Value.ListValues[0].Type != Atom {
		t.Errorf("Expected first element in cons'd list to have a type of atom. Received %d", c.(SExpression).Value.ListValues[0].Type)
	}

	if c.(SExpression).Value.ListValues[0].Value != "a" {
		t.Errorf("Expected first element in cons'd list to have a value of 'a'. Received %s", c.(SExpression).Value.ListValues[0].Value)
	}
}

func TestCons_OneAtomOneList(t *testing.T) {
	a := CreateSExpression(CreateAtom("a"))
	bList := CreateList()
	bList.ListValues = append(bList.ListValues, CreateAtom("b"))
	bList.ListValues = append(bList.ListValues, CreateAtom("c"))
	b := CreateSExpression(bList)
	c, err := cons([]Token{a, b})
	if err != nil {
		t.Errorf("did not expect error")
	}

	if c.(SExpression).Value.Type != List {
		t.Errorf("Expected a list type. Received %d", c.(SExpression).Value.Type)
	}

	if len(c.(SExpression).Value.ListValues) != 2 {
		t.Errorf("Expected cons'd list to have length of 2. Received %d", len(c.(SExpression).Value.ListValues))
	}

	if c.(SExpression).Value.ListValues[0].Type != Atom {
		t.Errorf("Expected first element in cons'd list to have a type of atom. Received %d", c.(SExpression).Value.ListValues[0].Type)
	}

	if c.(SExpression).Value.ListValues[0].Value != "a" {
		t.Errorf("Expected first element in cons'd list to have a value of 'a'. Received %s", c.(SExpression).Value.ListValues[0].Value)
	}
}

func TestCons_TwoLists(t *testing.T) {
	aList := CreateList()
	aList.ListValues = append(aList.ListValues, CreateAtom("a"))
	aList.ListValues = append(aList.ListValues, CreateAtom("b"))
	a := CreateSExpression(aList)
	bList := CreateList()
	bList.ListValues = append(bList.ListValues, CreateAtom("c"))
	bList.ListValues = append(bList.ListValues, CreateAtom("d"))
	b := CreateSExpression(bList)
	c, err := cons([]Token{a, b})
	if err != nil {
		t.Errorf("did not expect error")
	}

	if c.(SExpression).Value.Type != List {
		t.Errorf("Expected a list type. Received %d", c.(SExpression).Value.Type)
	}

	if len(c.(SExpression).Value.ListValues) != 2 {
		t.Errorf("Expected cons'd list to have length of 2. Received %d", len(c.(SExpression).Value.ListValues))
	}

	if c.(SExpression).Value.ListValues[0].Type != List {
		t.Errorf("Expected first element in cons'd list to have a type of list. Received %d", c.(SExpression).Value.ListValues[0].Type)
	}

	if c.(SExpression).Value.ListValues[0].ListValues[0].Value != "a" {
		t.Errorf("Expected first element in cons'd list's list to have a value of 'a'. Received %s", c.(SExpression).Value.ListValues[0].ListValues[0].Value)
	}
}

func TestCar(t *testing.T) {
	aList := CreateList()
	aList.ListValues = append(aList.ListValues, CreateAtom("a"))
	aList.ListValues = append(aList.ListValues, CreateAtom("b"))
	a := CreateSExpression(aList)

	c, err := car([]Token{a})
	if err != nil {
		t.Errorf("did not expect error")
	}

	d, err := atom([]Token{c})
	if err != nil {
		t.Errorf("did not expect error")
	}

	if d.(SExpression).Value.Value == "T" && c.(SExpression).Value.Value != "a" {
		t.Errorf("Expected car(a b) to be a")
	}
}

func TestCdr(t *testing.T) {
	aList := CreateList()
	aList.ListValues = append(aList.ListValues, CreateAtom("a"))
	aList.ListValues = append(aList.ListValues, CreateAtom("b"))
	a := CreateSExpression(aList)

	c, err := cdr([]Token{a})
	if err != nil {
		t.Errorf("did not expect error")
	}

	d, err := atom([]Token{c})
	if err != nil {
		t.Errorf("did not expect error")
	}

	if d.(SExpression).Value.Value == "T" && len(c.(SExpression).Value.ListValues) != 1 && c.(SExpression).Value.ListValues[0].Value != "b" {
		t.Errorf("Expected cdr(a b) to be (b)")
	}
}

func TestEq(t *testing.T) {
	a := CreateSExpression(CreateAtom("a"))
	b := CreateSExpression(CreateAtom("b"))
	dList := CreateList()
	dList.ListValues = append(dList.ListValues, CreateAtom("c"))
	dList.ListValues = append(dList.ListValues, CreateAtom("d"))
	d := CreateSExpression(dList)

	c, err := eq([]Token{a, b})
	if err != nil {
		t.Errorf("did not expect error")
	}

	if c.(SExpression).Value.Value == "T" {
		t.Errorf("expected a != b")
	}

	c, err = eq([]Token{a, a})
	if err != nil {
		t.Errorf("did not expect error")
	}

	if c.(SExpression).Value.Value != "T" {
		t.Errorf("expected a == a")
	}

	_, err = eq([]Token{a, d})
	if err == nil {
		t.Errorf("expected an error when calling eq on a list")
	}
}

func TestAtom(t *testing.T) {
	a := CreateSExpression(CreateAtom("a"))
	bList := CreateList()
	bList.ListValues = append(bList.ListValues, CreateAtom("c"))
	bList.ListValues = append(bList.ListValues, CreateAtom("d"))
	b := CreateSExpression(bList)

	c, err := atom([]Token{a})
	if err != nil {
		t.Errorf("did not expect an error when calling eq on an atom")
	}

	if c.(SExpression).Value.Value != "T" {
		t.Errorf("Expected atom(Atom) to be true")
	}

	c, err = atom([]Token{b})
	if err != nil {
		t.Errorf("did not expect an error when calling eq on an atom")
	}

	if c.(SExpression).Value.Value == "T" {
		t.Errorf("Expected atom(List) to be true")
	}
}
