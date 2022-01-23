package lexicons

type LParen struct{}

func (l LParen) GetType() string {
	return "LParen"
}
func (l LParen) GetValue() string {
	return "("
}
