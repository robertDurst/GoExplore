package lexicons

type Lexicon interface {
	GetType() string
	GetValue() string
}
