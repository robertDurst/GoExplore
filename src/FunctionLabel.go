package GoExplore

type FunctionLabel struct {
	Name Ident
	Fn   Token
}

func (fl FunctionLabel) GetType() string {
	return "FunctionLabel"
}

func (fl FunctionLabel) PrettyFormat() string {
	return "Some FunctionLabel"
}

func CreateFunctionLabel(name Ident, fn Token) FunctionLabel {
	return FunctionLabel{Name: name, Fn: fn}
}
