package GoExplore

type FunctionLabel struct {
	Name Token
	Fn   Token
}

func (fl FunctionLabel) GetType() string {
	return "FunctionLabel"
}

func CreateFunctionLabel(name Token, fn Token) FunctionLabel {
	return FunctionLabel{Name: name, Fn: fn}
}
