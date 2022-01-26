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
	InitialCheckAndParseToken(t, "Constant", "ATOM")
}

func TestListWithinListTokenizesToConstant(t *testing.T) {
	InitialCheckAndParseToken(t, "Constant", "(HELLO (WORLD5 ((BAR))) FOO)")
}

func TestSimpleVariableTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "Variable", "foo")
}

func TestFunctionLabelTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "FunctionLabel", "label[foo nons]")
}

func TestFunctionLabelWithInnerFunctionLabelTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "FunctionLabel", "label[foo label[foobar sandwich]]")
}

func TestFunctionAtSignTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "FunctionAtSign", "@[[foo bar baz] label[some cons]]")
}

func TestFunctionAtSignInFunctionAtSignTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "FunctionAtSign", "@[[foo bar baz] @[[zee car zoom pool] cons]]")
}

func TestManyFunctionsInsideFunctions(t *testing.T) {
	InitialCheckAndParseToken(t, "FunctionLabel", "label[foo label[foobar @[[a b c d e] label[cons bar]]]]")
}

func TestSimpleConditionalStatementTokenizes(t *testing.T) {

	InitialCheckAndParseToken(t, "ConditionalStatement", "[[foo~T]]")
}

func TestComplexConditionalStatementTokenizes(t *testing.T) {
	tk := InitialCheckAndParseToken(t, "ConditionalStatement",
		`
[
	[@[[foo bar] cons]~label[cons foo]]
	[bar~F]
	[T~[[foo~T]]]
]
`)

	conditionalStatement := tk.(ConditionalStatement)
	if len(conditionalStatement.ConditionalPairs) != 3 {
		t.Errorf("Expected 3 conditional pairs. Received %d.", len(conditionalStatement.ConditionalPairs))
	}

	conditionalPair := conditionalStatement.ConditionalPairs[0]
	if conditionalPair.GetType() != "ConditionalPair" {
		t.Errorf("Expected 1st conditional pair. Received %s.", conditionalPair.GetType())
	}

	predicateForm := conditionalPair.Predicate
	if predicateForm.GetType() != "FunctionAtSign" {
		t.Errorf("Expected 1st conditional pair predicate token to be of type FunctionAtSign. Received %s.", predicateForm.GetType())
	}

	resultForm := conditionalPair.Result
	if resultForm.GetType() != "FunctionLabel" {
		t.Errorf("Expected 1st conditional pair result token to be of type FunctionLabel. Received %s.", resultForm.GetType())
	}
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
