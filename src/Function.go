package GoExplore

type Function struct {
	Name string
	Args []Token
}

func (f Function) GetType() string {
	return "Function"
}

func CreateFunction(name string, args []Token) Function {
	return Function{Name: name, Args: args}
}
