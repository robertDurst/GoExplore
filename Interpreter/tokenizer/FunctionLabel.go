package tokenizer

type FunctionLabel struct {
	Name Identifier
	Fn   Function
}

func (fl FunctionLabel) GetType() string {
	return "FunctionLabel"
}

func CreateFunctionLabel(name Identifier, fn Function) FunctionLabel {
	return FunctionLabel{Name: name, Fn: fn}
}
