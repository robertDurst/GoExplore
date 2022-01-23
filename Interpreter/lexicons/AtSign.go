package lexicons

type AtSign struct{}

func (a AtSign) GetType() string {
	return "AtSign"
}
func (a AtSign) GetValue() string {
	return "@"
}
