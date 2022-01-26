package tokenizer

type FunctionIdentifier struct {
	Name Identifier
}

func (fi FunctionIdentifier) GetType() string {
	return "FunctionIdentifier"
}

func CreateFunctionIdentifier(name Identifier) FunctionIdentifier {
	return FunctionIdentifier{Name: name}
}
