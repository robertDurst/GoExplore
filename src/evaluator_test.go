package GoExplore

import (
	"testing"
)

func TestEvalForm_SExpression_Atom(t *testing.T) {
	input := "FOO"

	le := CreateLexarExecutor()
	ls, err := le.Lex(input)
	if err != nil {
		t.Errorf("did not expect a lexar error")
	}

	tk, err := Tokenize(ls)
	if err != nil {
		t.Errorf("did not expect an error")
	}

	finalTk := eval(tk)
	if finalTk.GetType() != "SExpression" {
		t.Errorf("Expected SExpression. Received %s.", finalTk.GetType())
	}

	sexp := finalTk.(SExpression)
	if sexp.Value.Value != "FOO" {
		t.Errorf("Expected FOO. Received %s.", sexp.Value.Value)
	}
}

func TestEvalForm_SExpression_List(t *testing.T) {
	input := "(FOO BAR)"

	le := CreateLexarExecutor()
	ls, err := le.Lex(input)
	if err != nil {
		t.Errorf("did not expect a lexar error")
	}

	tk, err := Tokenize(ls)
	if err != nil {
		t.Errorf("did not expect an error")
	}

	finalTk := eval(tk)
	if finalTk.GetType() != "SExpression" {
		t.Errorf("Expected SExpression. Received %s.", finalTk.GetType())
	}

	sexp := finalTk.(SExpression)
	if len(sexp.Value.ListValues) != 2 {
		t.Errorf("Expected list of length 2. Received length of %d.", len(sexp.Value.ListValues))
	}
}

func TestEvalForm_Variable(t *testing.T) {
	input := "foo"

	le := CreateLexarExecutor()
	ls, err := le.Lex(input)
	if err != nil {
		t.Errorf("did not expect a lexar error")
	}

	tk, err := Tokenize(ls)
	if err != nil {
		t.Errorf("did not expect an error")
	}

	finalTk := eval(tk)
	if finalTk.GetType() != "Variable" {
		t.Errorf("Expected SExpression. Received %s.", finalTk.GetType())
	}

	sexp := finalTk.(Variable)
	if sexp.Value != "foo" {
		t.Errorf("Expected foo. Received %s.", sexp.Value)
	}
}

func TestEvalForm_Var(t *testing.T) {
	input := "foo"

	le := CreateLexarExecutor()
	ls, err := le.Lex(input)
	if err != nil {
		t.Errorf("did not expect a lexar error")
	}

	tk, err := Tokenize(ls)
	if err != nil {
		t.Errorf("did not expect an error")
	}

	finalTk := eval(tk)
	if finalTk.GetType() != "Variable" {
		t.Errorf("Expected SExpression. Received %s.", finalTk.GetType())
	}

	sexp := finalTk.(Variable)
	if sexp.Value != "foo" {
		t.Errorf("Expected foo. Received %s.", sexp.Value)
	}
}

func TestEvalForm_FunctionIdentifier_Simple(t *testing.T) {
	input := "cons[[(FOO BAR)] [BAZ]]"

	le := CreateLexarExecutor()
	ls, err := le.Lex(input)
	if err != nil {
		t.Errorf("did not expect a lexar error")
	}

	tk, err := Tokenize(ls)
	if err != nil {
		t.Errorf("did not expect an error")
	}

	finalTk := eval(tk)
	if finalTk.GetType() != "SExpression" {
		t.Errorf("Expected SExpression. Received %s.", finalTk.GetType())
	}

	sexp := finalTk.(SExpression)
	if len(sexp.Value.ListValues) != 2 {
		t.Errorf("Expected list of length 2. Received length of %d.", len(sexp.Value.ListValues))
	}

	if len(sexp.Value.ListValues[0].ListValues) != 2 {
		t.Errorf("Expected first element to be a list of length 2. Received length of %d.", len(sexp.Value.ListValues[0].ListValues))
	}

	if sexp.Value.ListValues[0].ListValues[0].Value != "FOO" {
		t.Errorf("Expected FOO. Received %s.", sexp.Value.ListValues[0].ListValues[0].Value)
	}
}

func TestEvalForm_FunctionIdentifier_Simple2(t *testing.T) {
	input := `
cons[
	[
		cdr[[(FOO BAR)]]
	] 
	[BAZ]
]
`

	le := CreateLexarExecutor()
	ls, err := le.Lex(input)
	if err != nil {
		t.Errorf("did not expect a lexar error")
	}

	tk, err := Tokenize(ls)
	if err != nil {
		t.Errorf("did not expect an error, %s", err.Error())
	}

	finalTk := eval(tk)
	if finalTk.GetType() != "SExpression" {
		t.Errorf("Expected SExpression. Received %s.", finalTk.GetType())
	}

	sexp := finalTk.(SExpression)
	if len(sexp.Value.ListValues) != 2 {
		t.Errorf("Expected list of length 2. Received length of %d.", len(sexp.Value.ListValues))
	}

	if len(sexp.Value.ListValues[0].ListValues) != 1 {
		t.Errorf("Expected first element to be a list of length 1. Received length of %d.", len(sexp.Value.ListValues[0].ListValues))
	}

	if sexp.Value.ListValues[0].ListValues[0].Value != "BAR" {
		t.Errorf("Expected BAR. Received %s.", sexp.Value.ListValues[0].ListValues[0].Value)
	}
}

func TestEvalForm_FunctionIdentifier_Simple3(t *testing.T) {
	input := `
cons[
		[
			cdr[
					[
						cons[
							[(FOO BAR)]
							[(BAZ FAZ)]
						]
					]
			]
		] 
		[
			car[
				[(GAZ BAZ)]
			]
		]
]
`

	le := CreateLexarExecutor()
	ls, err := le.Lex(input)
	if err != nil {
		t.Errorf("did not expect a lexar error, %s", err.Error())
	}

	tk, err := Tokenize(ls)
	if err != nil {
		t.Errorf("did not expect an error, %s", err.Error())
	}

	finalTk := eval(tk)
	if finalTk.GetType() != "SExpression" {
		t.Errorf("Expected SExpression. Received %s.", finalTk.GetType())
	}

	sexp := finalTk.(SExpression)
	if len(sexp.Value.ListValues) != 2 {
		t.Errorf("Expected list of length 2. Received length of %d.", len(sexp.Value.ListValues))
	}

	if len(sexp.Value.ListValues[0].ListValues) != 1 {
		t.Errorf("Expected first element to be a list of length 1. Received length of %d.", len(sexp.Value.ListValues[0].ListValues))
	}

	// make sure we have ((BAZ FAZ) GAZ)
	if sexp.Value.ListValues[0].ListValues[0].ListValues[0].Value != "BAZ" {
		t.Errorf("Expected BAZ. Received %s.", sexp.Value.ListValues[0].ListValues[0].ListValues[0].Value)
	}
	if sexp.Value.ListValues[0].ListValues[0].ListValues[1].Value != "FAZ" {
		t.Errorf("Expected FAZ. Received %s.", sexp.Value.ListValues[0].ListValues[0].ListValues[1].Value)
	}
	if sexp.Value.ListValues[1].Value != "GAZ" {
		t.Errorf("Expected GAZ. Received %s.", sexp.Value.ListValues[1].Value)
	}
}
