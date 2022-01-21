package interpreter

func evaluateAtom(atom Atom, rest []SExpression) SExpression {
	return evaluate(rest)
}

func evaluateList(list List, rest []SExpression) SExpression {
	return evaluate(rest)
}

func evaluate(tokens []SExpression) SExpression {
	cur := tokens[0]
	if len(tokens) == 1 {
		return cur
	}

	rest := tokens[1:]
	if cur.GetName() == "List" {
		field, _ := cur.(List)
		return evaluateList(field, rest)
	} else if cur.GetName() == "Atom" {
		field, _ := cur.(Atom)
		return evaluateAtom(Atom(field), rest)
	}

	panic("Unexpected token type received during evaluation")
}
