package lexicons

type Dot struct{}

func (d Dot) GetType() string {
	return "Dot"
}
func (d Dot) GetValue() string {
	return "."
}
