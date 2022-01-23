package lexicons

type Atom struct {
	value string
}

func (a Atom) GetType() string {
	return "Atom"
}
func (a Atom) GetValue() string {
	return a.value
}

func CreateAtom(value string) Atom {
	return Atom{value: value}
}
