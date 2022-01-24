package lexicons

type List struct {
	Value []SExpression
}

func (l List) GetType() string {
	return "List"
}

func (l List) GetValue() []SExpression {
	return l.Value
}

func CreateList() List {
	return List{Value: make([]SExpression, 0)}
}
