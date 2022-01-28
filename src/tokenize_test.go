package GoExplore

import (
	"testing"
)

func TestEmptyInputTokenizer(t *testing.T) {
	le := CreateLexarExecutor()
	ls, err := le.Lex("")
	if err != nil {
		t.Errorf("did not expect a lexar error")
	}

	_, err = Tokenize(ls)
	if err == nil {
		t.Errorf("expected an error")
	}
}

func TestSimpleConstantTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "SExpression", "ATOM")
}

func TestListWithinListTokenizesToConstant(t *testing.T) {
	InitialCheckAndParseToken(t, "SExpression", "(HELLO (WORLD5 ((BAR))) FOO)")
}

func TestSimpleConditionalStatementTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "ConditionalStatement", "[[foo~T]]")
}

func InitialCheckAndParseToken(t *testing.T, expectedTopLevelType string, input string) Token {
	le := CreateLexarExecutor()
	ls, err := le.Lex(input)
	if err != nil {
		t.Errorf("did not expect a lexar error")
	}

	tk, err := Tokenize(ls)
	if err != nil {
		t.Errorf("did not expect an error")
	}

	if tk.GetType() != expectedTopLevelType {
		t.Errorf("Expected outer token to be of type %s. Received %s.", expectedTopLevelType, tk.GetType())
	}

	return tk
}
