package lexicons

type RParen struct{}

func (r RParen) GetType() string {
	return "RParen"
}
func (r RParen) GetValue() string {
	return ")"
}
