package lexicons

type RSquareBracket struct{}

func (r RSquareBracket) GetType() string {
	return "RSquareBracket"
}
func (r RSquareBracket) GetValue() string {
	return "]"
}
