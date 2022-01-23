package lexicons

type Squiggle struct{}

func (s Squiggle) GetType() string {
	return "Squiggle"
}
func (s Squiggle) GetValue() string {
	return "~"
}
