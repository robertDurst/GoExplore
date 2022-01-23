package lexicons

type Semicolon struct{}

func (s Semicolon) GetType() string {
	return "Semicolon"
}
func (s Semicolon) GetValue() string {
	return ";"
}
