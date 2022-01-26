package lexicons

type ArgList struct {
	Value []Lexicon
}

func (a ArgList) GetType() string {
	return "ArgList"
}

func (a ArgList) GetValue() []Lexicon {
	return a.Value
}

func CreateArgList() ArgList {
	return ArgList{Value: make([]Lexicon, 0)}
}
