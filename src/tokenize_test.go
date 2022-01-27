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

func TestSimpleVariableTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "Variable", "foo")
}

func TestFunctionLabelTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "FunctionLabel", "label[foo nons]")
}

func TestFunctionLabelWithInnerFunctionLabelTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "FunctionLabel", `
label[foo 
	label[foobar sandwich]
]
`)
}

func TestSimpleConditionalStatementTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "ConditionalStatement", "[[foo~T]]")
}

func TestComplexConditionalStatementTokenizes(t *testing.T) {
	tk := InitialCheckAndParseToken(t, "ConditionalStatement",
		`
[
	[
		cons[[arg1] [arg2]]
		~
		label[foz foo]
	]
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
	if predicateForm.GetType() != "Function" {
		t.Errorf("Expected 1st conditional pair predicate token to be of type Function. Received %s.", predicateForm.GetType())
	}
	predicateFormFunction := predicateForm.(Function)
	predicateFormFunctionName := predicateFormFunction.Name
	if predicateFormFunctionName != "cons" {
		t.Errorf("Expected first variable in first conditional pair predicate to be cons. Received %s.", predicateFormFunctionName)
	}

	predicateFormFunctionArgs := predicateFormFunction.Args
	if len(predicateFormFunctionArgs) != 2 {
		t.Errorf("Expected predicate in first conditional pair to have 2 args. Received %d.", len(predicateFormFunctionArgs))
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
