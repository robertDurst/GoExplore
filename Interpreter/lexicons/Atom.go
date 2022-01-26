package lexicons

type Atom struct {
	value         string
	isSExpression bool
}

func (a Atom) GetType() string {
	return "Atom"
}
func (a Atom) GetValue() string {
	return a.value
}

func (a Atom) IsSExpression() bool {
	return a.isSExpression
}

func CreateAtom(value string, isSExpression bool) Atom {
	return Atom{value: value, isSExpression: isSExpression}
}
