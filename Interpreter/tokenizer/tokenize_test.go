package tokenizer

import (
	"GoExplore/interpreter/lexar"
	"testing"
)

func TestEmptyInputTokenizer(t *testing.T) {
	le := lexar.CreateLexarExecutor()
	ls, err := le.Lex("")
	if err != nil {
		t.Errorf("did not expect a lexar error")
	}

	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 0 {
		t.Errorf("Expected 0 tokens from empty input. Received %d.", len(tks))
	}
}

func TestSimpleConstantTokenizes(t *testing.T) {
	InitialCheckAndParseForm(t, "Constant", "ATOM")
}

func TestListWithinListTokenizesToConstant(t *testing.T) {
	InitialCheckAndParseForm(t, "Constant", "(HELLO (WORLD5 ((BAR))) FOO)")
}

func TestSimpleVariableTokenizes(t *testing.T) {
	InitialCheckAndParseForm(t, "Variable", "foo")
}

func TestFunctionIdentifierTokenizes(t *testing.T) {
	InitialCheckAndParseForm(t, "FunctionIdentifier", "atom")
}

func TestFunctionLabelTokenizes(t *testing.T) {
	InitialCheckAndParseForm(t, "FunctionLabel", "label[foo;cons]")
}

func TestFunctionLabelWithInnerFunctionLabelTokenizes(t *testing.T) {
	InitialCheckAndParseForm(t, "FunctionLabel", "label[foo;label[foobar;sandwich]]")
}

func TestFunctionAtSignTokenizes(t *testing.T) {
	InitialCheckAndParseForm(t, "FunctionAtSign", "@[[foo;bar;baz];label[some;cons]]")
}

func TestFunctionAtSignInFunctionAtSignTokenizes(t *testing.T) {
	InitialCheckAndParseForm(t, "FunctionAtSign", "@[[foo;bar;baz];@[[zee;car;zoom;pool];cons]]")
}

func TestManyFunctionsInsideFunctions(t *testing.T) {
	InitialCheckAndParseForm(t, "FunctionLabel", "label[foo;label[foobar;@[[a;b;c;d;e];label[cons;bar]]]]")
}

func TestSimpleConditionalStatementTokenizes(t *testing.T) {

	InitialCheckAndParseForm(t, "ConditionalStatement", "[foo~T]")
}

func TestComplexConditionalStatementTokenizes(t *testing.T) {
	form := InitialCheckAndParseForm(t, "ConditionalStatement", "[@[[foo;bar];cons]~label[cons;foo];bar~F;T~[foo~t;]]")

	conditionalStatement := form.Value.(ConditionalStatement)
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

func InitialCheckAndParseForm(t *testing.T, expectedTopLevelType string, input string) Form {
	le := lexar.CreateLexarExecutor()
	ls, err := le.Lex(input)
	if err != nil {
		t.Errorf("did not expect a lexar error")
	}

	tokenizer := CreateTokenizer(ls)
	tks := tokenizer.tokenize()

	if len(tks) != 1 {
		t.Errorf("Expected 1 tokens from function label. Received %d.", len(tks))
	}

	if tks[0].GetType() != "Form" {
		t.Errorf("Expected outer token to be of type Form. Received %s.", tks[0].GetType())
	}

	form := tks[0].(Form)
	if form.Value.GetType() != expectedTopLevelType {
		t.Errorf("Expected form's inner value to be a FunctionLabel. Received %s.", form.Value.GetType())
	}

	return form
}
