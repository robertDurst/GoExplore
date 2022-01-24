package interpreter

import (
	"GoExplore/interpreter/lexicons"
	"GoExplore/interpreter/tokens"
)

type Tokenizer struct {
	lexs []lexicons.Lexicon
	tks  []tokens.Token
	i    int
}

func CreateTokenizer(lexs []lexicons.Lexicon) Tokenizer {
	return Tokenizer{lexs: lexs, tks: make([]tokens.Token, 0), i: 0}
}

func (t Tokenizer) GetTokens() []tokens.Token {
	return t.tks
}

func (t Tokenizer) tokenize() []tokens.Token {

	for t.i < len(t.lexs) {
		t.parseForm()
	}

	return t.tks
}

func (t *Tokenizer) parseForm() {
	cur := t.lexs[t.i]

	switch cur.GetType() {
	case "Atom":
		atom := cur.(lexicons.Atom)

		if atom.IsSExpression() {
			constant := tokens.CreateConstant(atom)
			form := tokens.CreateForm(constant)
			t.tks = append(t.tks, form)

			t.i++
			return
		}

		identifier := tokens.CreateIdentifier(atom.GetValue())
		variable := tokens.CreateVariable(identifier)
		form := tokens.CreateForm(variable)
		t.tks = append(t.tks, form)

		t.i++
		return
	}
}
