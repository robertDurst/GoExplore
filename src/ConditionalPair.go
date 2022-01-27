package GoExplore

type ConditionalPair struct {
	Predicate Token
	Result    Token
}

func (cp ConditionalPair) GetType() string {
	return "ConditionalPair"
}

func (cp ConditionalPair) PrettyFormat() string {
	return "Some ConditionalPair"
}

func CreateConditionalPair(predicate Token, result Token) ConditionalPair {
	return ConditionalPair{Predicate: predicate, Result: result}
}
