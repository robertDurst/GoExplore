package tokenizer

type FunctionAtSign struct {
	Vars []Variable
	Rest Form
}

func (fas FunctionAtSign) GetType() string {
	return "FunctionAtSign"
}

func CreateFunctionAtSign(vars []Variable, rest Form) FunctionAtSign {
	return FunctionAtSign{Vars: vars, Rest: rest}
}
