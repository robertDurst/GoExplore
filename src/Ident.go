package GoExplore

type Ident struct {
	Value string
}

func (i Ident) GetType() string {
	return "Ident"
}

func (i Ident) PrettyFormat() string {
	return i.Value
}

func CreateIdent(value string) Ident {
	return Ident{Value: value}
}
