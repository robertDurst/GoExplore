package GoExplore

import "fmt"

type Evaluator struct {
	functionMap map[string]func([]Token) (Token, error)
}

func CreateEvaluator() Evaluator {
	fmap := make(map[string]func([]Token) (Token, error))
	fmap["cons"] = cons
	fmap["car"] = car
	fmap["cdr"] = cdr
	fmap["eq"] = eq
	fmap["atom"] = atom

	return Evaluator{functionMap: fmap}
}

func (e Evaluator) eval(tk Token) Token {
	switch tk.GetType() {
	case "SExpression":
		return e.evalSExpression(tk.(SExpression))
	case "Variable":
		return e.evalVariable(tk.(Variable))
	case "Function":
		return e.evalFunction(tk.(Function))
	case "FunctionLabel":
		return e.evalFunctionLabel(tk.(FunctionLabel))
	case "ConditionalStatement":
		return e.evalConditionalStatement(tk.(ConditionalStatement))
	default:
		return nil
	}
}

func (e Evaluator) evalSExpression(sexp SExpression) Token {
	return sexp
}

func (e Evaluator) evalVariable(v Variable) Token {
	return v
}

func (e Evaluator) evalFunction(fn Function) Token {
	args := make([]Token, 0)
	for _, arg := range fn.Args {
		evaldArg := e.eval(arg)
		if evaldArg.GetType() != "SExpression" {
			panic("expected only SExpression args in FunctionIdentifier")
		}
		args = append(args, evaldArg)
	}

	name := fn.Name
	if f, ok := e.functionMap[name]; ok {
		val, _ := f(args)
		return val
	}
	return nil
}

// returns whether or not was able to successfully "create" function
func (e Evaluator) evalFunctionLabel(fl FunctionLabel) Token {
	name := fl.Name

	fmt.Printf("Creating function %s.", name)

	return boolToSExpression(true)
}

func (e Evaluator) evalConditionalStatement(cs ConditionalStatement) Token {
	return nil
}
