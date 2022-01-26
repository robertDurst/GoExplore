package GoExplore

import (
	"testing"
)

func TestCons_TwoAtoms(t *testing.T) {
	a := CreateSExpression(CreateAtom("a"))
	b := CreateSExpression(CreateAtom("b"))
	c := cons(a, b)

	if c.Value.Type != List {
		t.Errorf("Expected a list type. Received %d", c.Value.Type)
	}

	if len(c.Value.ListValues) != 2 {
		t.Errorf("Expected cons'd list to have length of 2. Received %d", len(c.Value.ListValues))
	}

	if c.Value.ListValues[0].Type != Atom {
		t.Errorf("Expected first element in cons'd list to have a type of atom. Received %d", c.Value.ListValues[0].Type)
	}

	if c.Value.ListValues[0].Value != "a" {
		t.Errorf("Expected first element in cons'd list to have a value of 'a'. Received %s", c.Value.ListValues[0].Value)
	}
}

func TestCons_OneAtomOneList(t *testing.T) {
	a := CreateSExpression(CreateAtom("a"))
	bList := CreateList()
	bList.ListValues = append(bList.ListValues, CreateAtom("b"))
	bList.ListValues = append(bList.ListValues, CreateAtom("c"))
	b := CreateSExpression(bList)
	c := cons(a, b)

	if c.Value.Type != List {
		t.Errorf("Expected a list type. Received %d", c.Value.Type)
	}

	if len(c.Value.ListValues) != 2 {
		t.Errorf("Expected cons'd list to have length of 2. Received %d", len(c.Value.ListValues))
	}

	if c.Value.ListValues[0].Type != Atom {
		t.Errorf("Expected first element in cons'd list to have a type of atom. Received %d", c.Value.ListValues[0].Type)
	}

	if c.Value.ListValues[0].Value != "a" {
		t.Errorf("Expected first element in cons'd list to have a value of 'a'. Received %s", c.Value.ListValues[0].Value)
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
	c := cons(a, b)

	if c.Value.Type != List {
		t.Errorf("Expected a list type. Received %d", c.Value.Type)
	}

	if len(c.Value.ListValues) != 2 {
		t.Errorf("Expected cons'd list to have length of 2. Received %d", len(c.Value.ListValues))
	}

	if c.Value.ListValues[0].Type != List {
		t.Errorf("Expected first element in cons'd list to have a type of list. Received %d", c.Value.ListValues[0].Type)
	}

	if c.Value.ListValues[0].ListValues[0].Value != "a" {
		t.Errorf("Expected first element in cons'd list's list to have a value of 'a'. Received %s", c.Value.ListValues[0].ListValues[0].Value)
	}
}

func TestCar(t *testing.T) {
	aList := CreateList()
	aList.ListValues = append(aList.ListValues, CreateAtom("a"))
	aList.ListValues = append(aList.ListValues, CreateAtom("b"))
	a := CreateSExpression(aList)

	c, err := car(a)
	if err != nil {
		t.Errorf("did not expect error")
	}

	if !atom(c) && c.Value.Value != "a" {
		t.Errorf("Expected car(a b) to be a")
	}
}

func TestCdr(t *testing.T) {
	aList := CreateList()
	aList.ListValues = append(aList.ListValues, CreateAtom("a"))
	aList.ListValues = append(aList.ListValues, CreateAtom("b"))
	a := CreateSExpression(aList)

	c, err := cdr(a)
	if err != nil {
		t.Errorf("did not expect error")
	}

	if atom(c) && len(c.Value.ListValues) != 1 && c.Value.ListValues[0].Value != "b" {
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

	c, err := eq(a, b)
	if err != nil {
		t.Errorf("did not expect error")
	}

	if c {
		t.Errorf("expected a != b")
	}

	c, err = eq(a, a)
	if err != nil {
		t.Errorf("did not expect error")
	}

	if !c {
		t.Errorf("expected a == a")
	}

	_, err = eq(a, d)
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

	if !atom(a) {
		t.Errorf("Expected atom(Atom) to be true")
	}

	if atom(b) {
		t.Errorf("Expected atom(List) to be true")
	}
}
