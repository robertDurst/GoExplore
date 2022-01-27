package GoExplore

import "errors"

// elementary SExpression functions
func cons(a SExpression, b SExpression) SExpression {
	encompassingList := CreateList()
	encompassingList.ListValues = append(encompassingList.ListValues, a.Value)
	encompassingList.ListValues = append(encompassingList.ListValues, b.Value)
	return CreateSExpression(encompassingList)
}

func car(a SExpression) (SExpression, error) {
	if a.Value.Type != List || len(a.Value.ListValues) == 0 {
		return CreateSExpression(CreateAtom("")), errors.New("expected a list with at least one value")
	}

	return CreateSExpression(a.Value.ListValues[0]), nil
}

func cdr(a SExpression) (SExpression, error) {
	if a.Value.Type != List || len(a.Value.ListValues) == 0 {
		return CreateSExpression(CreateAtom("")), errors.New("expected a list with at least one value")
	}

	list := CreateList()
	list.ListValues = append(list.ListValues, a.Value.ListValues[1:]...)

	return CreateSExpression(list), nil
}

func eq(a SExpression, b SExpression) (bool, error) {
	if !atom(a) || !atom(b) {
		return false, errors.New("expected both arguments to be of type Atom")
	}

	return a.Value.Value == b.Value.Value, nil
}

func atom(a SExpression) bool {
	return a.Value.Type == Atom
}

// auxillary functions for evalquote
// func equal() {
// }

// func subst() {
// }

// func null() {
// }

// // func append() {
// // }

// func member() {
// }

// func pairlis() {
// }

// func assoc() {
// }

// func sublis() {
// }

func eval(tk Token) Token {
	switch tk.GetType() {
	case "SExpression":
		return evalSExpression(tk.(SExpression))
	case "Variable":
		return evalVariable(tk.(Variable))
	case "Function":
		return evalFunction(tk.(Function))
	case "ConditionalStatement":
		return evalConditionalStatement(tk.(ConditionalStatement))
	default:
		return nil
	}
}

func evalSExpression(sexp SExpression) Token {
	return sexp
}

func evalVariable(v Variable) Token {
	return v
}

func evalFunction(fn Function) Token {
	switch fn.GetType() {
	case "FunctionAtSign":
	case "Label":
	default:
		return nil
	}
	return nil
}

func evalFunctionAtSign(fn FunctionAtSign) Token {
	return nil
}

func evalFunctionLabel(fn FunctionLabel) Token {
	return nil
}

func evalConditionalStatement(cs ConditionalStatement) Token {
	return nil
}
