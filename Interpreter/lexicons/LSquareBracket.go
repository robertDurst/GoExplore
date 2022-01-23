package lexicons

type LSquareBracket struct{}

func (l LSquareBracket) GetType() string {
	return "LSquareBracket"
}
func (l LSquareBracket) GetValue() string {
	return "["
}
