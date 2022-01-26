package lexicons

type Lexicon interface {
	GetType() string
}

type SExpression interface {
	GetType() string
}
