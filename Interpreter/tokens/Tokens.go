package tokens

import "GoExplore/interpreter/lexicons"

type Token interface {
	GetType() string
}

type Form struct {
	Value Token
}

func (f Form) GetType() string {
	return "Form"
}

func CreateForm(token Token) Form {
	return Form{Value: token}
}

type Constant struct {
	Value lexicons.SExpression
}

func (c Constant) GetType() string {
	return "Constant"
}

func CreateConstant(value lexicons.SExpression) Constant {
	return Constant{Value: value}
}

type Variable struct {
	Value Identifier
}

func (v Variable) GetType() string {
	return "Variable"
}

func CreateVariable(value Identifier) Variable {
	return Variable{Value: value}
}

type Identifier struct {
	Value string
}

func (i Identifier) GetType() string {
	return "Identifier"
}

func CreateIdentifier(value string) Identifier {
	return Identifier{Value: value}
}
