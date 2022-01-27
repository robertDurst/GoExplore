package GoExplore

type ConditionalStatement struct {
	ConditionalPairs []ConditionalPair
}

func (cs ConditionalStatement) GetType() string {
	return "ConditionalStatement"
}

func (cs ConditionalStatement) PrettyFormat() string {
	return "Some ConditionalStatement"
}

func CreateConditionalStatement() ConditionalStatement {
	return ConditionalStatement{ConditionalPairs: make([]ConditionalPair, 0)}
}
