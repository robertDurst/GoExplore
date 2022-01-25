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

type Function interface {
	GetType() string
}

type FunctionIdentifier struct {
	Name Identifier
}

func (fi FunctionIdentifier) GetType() string {
	return "FunctionIdentifier"
}

func CreateFunctionIdentifier(name Identifier) FunctionIdentifier {
	return FunctionIdentifier{Name: name}
}

type FunctionLabel struct {
	Name Identifier
	Fn   Function
}

func (fl FunctionLabel) GetType() string {
	return "FunctionLabel"
}

func CreateFunctionLabel(name Identifier, fn Function) FunctionLabel {
	return FunctionLabel{Name: name, Fn: fn}
}

type VarList struct {
	Value []Token
}

func (v VarList) GetType() string {
	return "VarList"
}

func CreateVarList() VarList {
	return VarList{Value: make([]Token, 0)}
}

type FunctionAtSign struct {
	Vars VarList
	Rest Form
}

func (fas FunctionAtSign) GetType() string {
	return "FunctionAtSign"
}

func CreateFunctionAtSign(vars VarList, rest Form) FunctionAtSign {
	return FunctionAtSign{Vars: vars, Rest: rest}
}
