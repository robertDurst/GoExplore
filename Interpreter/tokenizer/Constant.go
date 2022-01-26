package tokenizer

import "GoExplore/interpreter/lexar"

type Constant struct {
	Value lexar.Lexicon
}

func (c Constant) GetType() string {
	return "Constant"
}

func CreateConstant(value lexar.Lexicon) Constant {
	return Constant{Value: value}
}
