package GoExplore

type FunctionIdentifier struct {
	Name Variable
}

func (fi FunctionIdentifier) GetType() string {
	return "FunctionIdentifier"
}

func CreateFunctionIdentifier(name Variable) FunctionIdentifier {
	return FunctionIdentifier{Name: name}
}
