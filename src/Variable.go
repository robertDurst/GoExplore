package GoExplore

type Variable struct {
	Value Identifier
}

func (v Variable) GetType() string {
	return "Variable"
}

func CreateVariable(value Identifier) Variable {
	return Variable{Value: value}
}
