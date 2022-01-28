package GoExplore

import "fmt"

type Evaluator struct {
	functionMap map[string]func([]Token) (Token, error)
	labelMap    map[string]Function
}

func CreateEvaluator() Evaluator {
	fmap := make(map[string]func([]Token) (Token, error))
	fmap["cons"] = cons
	fmap["car"] = car
	fmap["cdr"] = cdr
	fmap["eq"] = eq
	fmap["atom"] = atom

	lmap := make(map[string]Function)

	return Evaluator{functionMap: fmap, labelMap: lmap}
}

func (e Evaluator) Eval(tk Token) Token {
	switch tk.GetType() {
	case "SExpression":
		return e.evalSExpression(tk.(SExpression))
	case "Ident":
		return e.evalIdent(tk.(Ident))
	case "Function":
		return e.evalFunction(tk.(Function))
	case "ConditionalStatement":
		return e.evalConditionalStatement(tk.(ConditionalStatement))
	default:
		return nil
	}
}

func (e Evaluator) evalSExpression(sexp SExpression) Token {
	return sexp
}

func (e Evaluator) evalIdent(ident Ident) Token {
	return ident
}

func (e Evaluator) evalFunction(fn Function) Token {
	args := make([]Token, 0)
	for _, arg := range fn.Args {
		evaldArg := e.Eval(arg)
		if evaldArg.GetType() == "SExpression" || evaldArg.GetType() == "Ident" {
			args = append(args, evaldArg)
		} else {
			panic("Function arguments must be either a SExpression or a Variable")
		}
	}

	name := fn.Name
	if f, ok := e.functionMap[name]; ok {
		val, _ := f(args)
		return val
	}

	if f, ok := e.labelMap[name]; ok {
		fmt.Printf("We want to execute %s() which has %d args and we've supplied %d args\n", f.Name, len(f.Args), len(fn.Args))
		return nil
	}

	return nil
}

func (e Evaluator) evalConditionalStatement(cs ConditionalStatement) Token {
	for _, cp := range cs.ConditionalPairs {
		isTruthy := e.Eval(cp.Predicate)

		if isTruthy.GetType() == "SExpression" && isTruthy.(SExpression).Value.Value == "T" {
			return e.Eval(cp.Result)
		}
	}

	return nil
}
