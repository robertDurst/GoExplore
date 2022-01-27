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
	InitialCheckAndParseToken(t, "FunctionContainer", "label[foo nons][[arg1] [arg2]]")
}

func TestFunctionLabelWithInnerFunctionLabelTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "FunctionContainer", `
label[foo 
	label[foobar sandwich][[arg1] [arg2]]
][[arg1] [arg2]]
`)
}

func TestFunctionAtSignTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "FunctionContainer", `
@[[foo bar baz] 
	label[some foz][[arg1] [arg2]]
][[arg3] [arg4]]
`)
}

func TestFunctionAtSignInFunctionAtSignTokenizes(t *testing.T) {
	InitialCheckAndParseToken(t, "FunctionContainer", `
@[[foo bar baz] 
	@[[zee car zoom pool] foz][[arg1] [arg2]]
][[arg3] [arg4]]
`)
}

func TestManyFunctionsInsideFunctions(t *testing.T) {
	InitialCheckAndParseToken(t, "FunctionContainer", `
label[foo 
	label[foobar 
		@[[a b c d e] 
			label[foz bar][[arg1] [arg2]]
		][[arg3] [arg4]]
	][[arg5] [arg6]]
][[arg7] [arg8]]
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
		@[[foo bar]foz][[arg1] [arg2]]
		~
		label[foz foo][[arg3] [arg4]]
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
	if predicateForm.GetType() != "FunctionContainer" {
		t.Errorf("Expected 1st conditional pair predicate token to be of type FunctionContainer. Received %s.", predicateForm.GetType())
	}
	predicateFormFunctionContainer := predicateForm.(FunctionContainer)
	predicateFormFunctionAtSign := predicateFormFunctionContainer.Fn.(FunctionAtSign)
	if predicateFormFunctionAtSign.Vars[0].(Variable).Value != "foo" {
		t.Errorf("Expected first variable in first conditional pair predicate to be foo. Received %s.", predicateFormFunctionAtSign.Vars[0].(Variable).Value)
	}

	predicateFormFunctionArgs := predicateFormFunctionContainer.Args
	if len(predicateFormFunctionArgs) != 2 {
		t.Errorf("Expected predicate in first conditional pair to have 2 args. Received %d.", len(predicateFormFunctionArgs))
	}

	resultForm := conditionalPair.Result
	if resultForm.GetType() != "FunctionContainer" {
		t.Errorf("Expected 1st conditional pair result token to be of type FunctionContainer. Received %s.", resultForm.GetType())
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
