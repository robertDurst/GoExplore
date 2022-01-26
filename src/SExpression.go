package GoExplore

type SExpression struct {
	Value Lexicon
}

func (s SExpression) GetType() string {
	return "SExpression"
}

func CreateSExpression(value Lexicon) SExpression {
	return SExpression{Value: value}
}
