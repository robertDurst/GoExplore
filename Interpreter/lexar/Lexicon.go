package lexar

const (
	ArgList = iota
	AtSign
	Atom
	Identifier
	List
	Semicolon
	Squiggle
	Unknown = 99
)

type Lexicon struct {
	IsSExpression bool
	ListValues    []Lexicon
	Type          int
	Value         string
}

func DefaultToken() Lexicon {
	return Lexicon{IsSExpression: false, ListValues: make([]Lexicon, 0), Type: Unknown, Value: ""}
}

func CreateArgList() Lexicon {
	lex := DefaultToken()
	lex.ListValues = make([]Lexicon, 0)
	lex.Type = ArgList
	return lex
}

func CreateAtom(value string) Lexicon {
	lex := DefaultToken()
	lex.IsSExpression = true
	lex.Type = Atom
	lex.Value = value
	return lex
}

func CreateAtSign() Lexicon {
	lex := DefaultToken()
	lex.Type = AtSign
	return lex
}

func CreateIdentifier(value string) Lexicon {
	lex := DefaultToken()
	lex.Value = value
	lex.Type = Identifier
	return lex
}

func CreateList() Lexicon {
	lex := DefaultToken()
	lex.IsSExpression = true
	lex.ListValues = make([]Lexicon, 0)
	lex.Type = List
	return lex
}

func CreateSemicolon() Lexicon {
	lex := DefaultToken()
	lex.Type = Semicolon
	return lex
}

func CreateSquiggle() Lexicon {
	lex := DefaultToken()
	lex.Type = Squiggle
	return lex
}
