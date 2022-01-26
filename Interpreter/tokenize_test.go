package interpreter

import (
	"GoExplore/interpreter/tokens"
	"testing"
)

func TestEmptyInputTokenizer(t *testing.T) {
	ls := lex("")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 0 {
		t.Errorf("Expected 0 tokens from empty input. Received %d.", len(tks))
	}
}

func TestSimpleConstantTokenizes(t *testing.T) {
	ls := lex("ATOM")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from simple atomic constant. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(tokens.Form)
	if form.Value.GetType() != "Constant" {
		t.Errorf("Expected form's inner value to be a Constant. Received %s.", form.Value.GetType())
	}
}

func TestListWithinListTokenizesToConstant(t *testing.T) {
	ls := lex("(HELLO (WORLD5 ((BAR))) FOO)")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from simple list constant. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(tokens.Form)
	if form.Value.GetType() != "Constant" {
		t.Errorf("Expected form's inner value to be a Constant. Received %s.", form.Value.GetType())
	}
}

func TestSimpleVariableTokenizes(t *testing.T) {
	ls := lex("foo")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from simple variable. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(tokens.Form)
	if form.Value.GetType() != "Variable" {
		t.Errorf("Expected form's inner value to be a Variable. Received %s.", form.Value.GetType())
	}
}

func TestFunctionIdentifierTokenizes(t *testing.T) {
	ls := lex("atom")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from function identifier. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(tokens.Form)
	if form.Value.GetType() != "FunctionIdentifier" {
		t.Errorf("Expected form's inner value to be a FunctionIdentifier. Received %s.", form.Value.GetType())
	}
}

func TestFunctionLabelTokenizes(t *testing.T) {
	ls := lex("label[foo;cons]")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from function label. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(tokens.Form)
	if form.Value.GetType() != "FunctionLabel" {
		t.Errorf("Expected form's inner value to be a FunctionLabel. Received %s.", form.Value.GetType())
	}
}

func TestFunctionLabelWithInnerFunctionLabelTokenizes(t *testing.T) {
	ls := lex("label[foo;label[foobar;sandwich]]")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from function label. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(tokens.Form)
	if form.Value.GetType() != "FunctionLabel" {
		t.Errorf("Expected form's inner value to be a FunctionLabel. Received %s.", form.Value.GetType())
	}
}

func TestFunctionAtSignTokenizes(t *testing.T) {
	ls := lex("@[[foo;bar;baz];label[some;cons]]")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from function label. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(tokens.Form)
	if form.Value.GetType() != "FunctionAtSign" {
		t.Errorf("Expected form's inner value to be a FunctionLabel. Received %s.", form.Value.GetType())
	}
}

func TestFunctionAtSignInFunctionAtSignTokenizes(t *testing.T) {
	ls := lex("@[[foo;bar;baz];@[[zee;car;zoom;pool];cons]]")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from function label. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(tokens.Form)
	if form.Value.GetType() != "FunctionAtSign" {
		t.Errorf("Expected form's inner value to be a FunctionLabel. Received %s.", form.Value.GetType())
	}
}

func TestManyFunctionsInsideFunctions(t *testing.T) {
	ls := lex("label[foo;label[foobar;@[[a;b;c;d;e];label[cons;bar]]]]")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from function label. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(tokens.Form)
	if form.Value.GetType() != "FunctionLabel" {
		t.Errorf("Expected form's inner value to be a FunctionLabel. Received %s.", form.Value.GetType())
	}
}

func TestSimpleConditionalStatementTokenizes(t *testing.T) {
	ls := lex("[foo~T]")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from function label. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(tokens.Form)
	if form.Value.GetType() != "ConditionalStatement" {
		t.Errorf("Expected form's inner value to be a FunctionLabel. Received %s.", form.Value.GetType())
	}
}

func TestComplexConditionalStatementTokenizes(t *testing.T) {
	ls := lex("[@[[foo;bar];cons]~label[cons;foo];bar~F;T~[foo~t;]]")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from function label. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(tokens.Form)
	if form.Value.GetType() != "ConditionalStatement" {
		t.Errorf("Expected form's inner value to be a FunctionLabel. Received %s.", form.Value.GetType())
	}

	conditionalStatement := form.Value.(tokens.ConditionalStatement)
	if len(conditionalStatement.ConditionalPairs) != 3 {
		t.Errorf("Expected 3 conditional pairs. Received %d.", len(conditionalStatement.ConditionalPairs))
	}

	conditionalPair := conditionalStatement.ConditionalPairs[0]
	if conditionalPair.GetType() != "ConditionalPair" {
		t.Errorf("Expected 1st conditional pair. Received %s.", conditionalPair.GetType())
	}
	if conditionalPair.Predicate.Value.GetType() != "FunctionAtSign" {
		t.Errorf("Expected 1st conditional pair predicate token to be of type FunctionAtSign. Received %s.", conditionalPair.Predicate.Value.GetType())
	}

	if conditionalPair.Result.Value.GetType() != "FunctionLabel" {
		t.Errorf("Expected 1st conditional pair result token to be of type FunctionLabel. Received %s.", conditionalPair.Result.Value.GetType())
	}
}
