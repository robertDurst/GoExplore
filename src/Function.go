package GoExplore

type Function struct {
	Name string
	Args []Token
}

func (f Function) GetType() string {
	return "Function"
}

func (f Function) PrettyFormat() string {
	return "Some Function"
}

func CreateFunction(name string, args []Token) Function {
	return Function{Name: name, Args: args}
}
