package tokenizer

type Form struct {
	Value Token
}

func (f Form) GetType() string {
	return "Form"
}

func CreateForm(token Token) Form {
	return Form{Value: token}
}
