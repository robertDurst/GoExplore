package GoExplore

type Constant struct {
	Value string
}

func (c Constant) GetType() string {
	return "Constant"
}

func CreateConstant(value string) Constant {
	return Constant{Value: value}
}
