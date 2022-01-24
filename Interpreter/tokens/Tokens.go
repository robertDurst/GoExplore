package tokens

type Token interface {
	GetType() string
}

type Constant struct{}
type Variable struct{}
type Function struct{}
type Form struct{}
type SExpression struct {
	Value string
}
