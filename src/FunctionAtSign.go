package GoExplore

type FunctionAtSign struct {
	Vars []Token
	Rest Token
}

func (fas FunctionAtSign) GetType() string {
	return "FunctionAtSign"
}

func CreateFunctionAtSign(vars []Token, rest Token) FunctionAtSign {
	return FunctionAtSign{Vars: vars, Rest: rest}
}
