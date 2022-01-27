package GoExplore

type FunctionContainer struct {
	Fn   Token
	Args []Token
}

func (fc FunctionContainer) GetType() string {
	return "FunctionContainer"
}

func CreateFunctionContainer(fn Token, args []Token) FunctionContainer {
	return FunctionContainer{Fn: fn, Args: args}
}
