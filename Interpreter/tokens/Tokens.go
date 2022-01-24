package tokens

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

type SExpression struct {
	Values []SubSExpression
}

func (se SExpression) GetType() string {
	return "SExpression"
}

func CreateSExpression(values []SubSExpression) SExpression {
	return SExpression{Values: values}
}

type SubSExpression interface {
	GetType() string
}

type List struct {
	Values []SubSExpression
}

func (l List) GetType() string {
	return "List"
}

type Atom struct {
	Value string
}

func (a Atom) GetType() string {
	return "Atom"
}
