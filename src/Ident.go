package GoExplore

type Ident struct {
	Value string
}

func (i Ident) GetType() string {
	return "Ident"
}

func CreateIdent(value string) Ident {
	return Ident{Value: value}
}
