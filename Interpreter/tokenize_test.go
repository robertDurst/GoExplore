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
		t.Errorf("Expected 1 tokens from empty input. Received %d.", len(tks))
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
	ls := lex("atom")
	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from empty input. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(tokens.Form)
	if form.Value.GetType() != "Variable" {
		t.Errorf("Expected form's inner value to be a Variable. Received %s.", form.Value.GetType())
	}
}
