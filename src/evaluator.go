package GoExplore

func eval(tk Token) Token {
	switch tk.GetType() {
	case "SExpression":
		return evalSExpression(tk.(SExpression))
	case "Variable":
		return evalVariable(tk.(Variable))
	case "FunctionContainer":
		return evalFunctionContainer(tk.(FunctionContainer))
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

func evalFunctionContainer(fc FunctionContainer) Token {
	switch fc.Fn.GetType() {
	case "FunctionIdentifier":
		return evalFunctionIdentifier(fc)
	case "FunctionAtSign":
		return nil
	case "FunctionLabel":
		return nil
	default:
		return nil
	}
}

func evalFunctionIdentifier(fc FunctionContainer) Token {
	args := make([]SExpression, 0)
	for _, arg := range fc.Args {
		evaldArg := eval(arg)
		if evaldArg.GetType() != "SExpression" {
			panic("expected only SExpression args in FunctionIdentifier")
		}
		args = append(args, evaldArg.(SExpression))
	}

	name := fc.Fn.(FunctionIdentifier).Name
	switch name.Value {
	case "cons":
		if len(args) != 2 {
			panic("expected 2 arg")
		}
		return cons(args[0], args[1])
	case "car":
		if len(args) != 1 {
			panic("expected 1 arg")
		}
		val, _ := car(args[0])
		return val
	case "cdr":
		if len(args) != 1 {
			panic("expected 1 arg")
		}
		val, _ := cdr(args[0])
		return val
	case "eq":
		if len(args) != 2 {
			panic("expected 2 arg")
		}
		return cons(args[0], args[1])
	case "atom":
		if len(args) != 1 {
			panic("expected 1 arg")
		}
		return boolToSExpression(atom(args[0]))
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

func boolToSExpression(condition bool) SExpression {
	if condition {
		return CreateSExpression(CreateAtom("T"))
	}
	return CreateSExpression(CreateAtom("F"))
}
