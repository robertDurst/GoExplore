package tokenizer

type Identifier struct {
	Value string
}

func (i Identifier) GetType() string {
	return "Identifier"
}

func CreateIdentifier(value string) Identifier {
	return Identifier{Value: value}
}
