package GoExplore

type Variable struct {
	Value string
}

func (v Variable) GetType() string {
	return "Variable"
}

func CreateVariable(value string) Variable {
	return Variable{Value: value}
}
