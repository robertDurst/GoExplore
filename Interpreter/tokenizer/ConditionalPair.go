package tokenizer

type ConditionalPair struct {
	Predicate Form
	Result    Form
}

func (cp ConditionalPair) GetType() string {
	return "ConditionalPair"
}

func CreateConditionalPair(predicate Form, result Form) ConditionalPair {
	return ConditionalPair{Predicate: predicate, Result: result}
}
